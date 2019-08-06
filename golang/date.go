package typesystem

import (
	"errors"
	"fmt"
	"time"
)

type DateType struct {
}

func NewDateType() iType {
	return &DateType{}
}

func (t *DateType) ConvToEls(key, operator, values string) (string, error) {
	switch operator {
	case Eq:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}

		return fmt.Sprintf(`{"bool": {"must": {"term": { %q: %q}}}}`, key, ds), nil
	case Ne:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must_not": {"term": {%q: %q}}}}`, key, ds), nil
	case Gt:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"gt": %q}}}}}`, key, ds), nil
	case Lt:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"lt": %q}}}}}`, key, ds), nil
	case Gte:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"gte": %q}}}}}`, key, ds), nil
	case Lte:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"lte": %q}}}}}`, key, ds), nil
	case InRange:
		fs := make([]string, 0)
		if err := parseJSON(values, &fs); err != nil {
			return "", err
		}
		if len(fs) < 2 {
			return "", errors.New("Worng format")
		}
		lower, upper := fs[0], fs[1]
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"gte": %q, "lte": %q}}}}}`, key, lower, upper), nil
	case NotInRange:
		fs := make([]string, 0)
		if err := parseJSON(values, &fs); err != nil {
			return "", err
		}
		if len(fs) < 2 {
			return "", errors.New("Wrong format")
		}
		lower, upper := fs[0], fs[1]
		return fmt.Sprintf(`{"bool": {"must_not": {"range": {%q: {"gte": %q, "lte": %q}}}}}`, key, lower, upper), nil
	case After:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"gt": %q}}}}}`, key, ds), nil
	case Before:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"lt": %q}}}}}`, key, ds), nil
	default:
		return "", fmt.Errorf("type/golang/datetime.go: unsupport operator, %v, %s, %s", key, operator, values)
	}
}

// values is in json format
func (t *DateType) Evaluate(obj interface{}, operator string, values string) (bool, error) {
	var object time.Time
	var err error
	if t, ok := obj.(time.Time); ok {
		object = t
	} else {
		sobj := fmt.Sprintf("%v", obj)
		if sobj != "" {
			object, err = time.Parse(time.RFC3339, sobj)
		}
	}

	switch operator {
	case Nad:
		return err != nil, nil
	case Ad:
		return err == nil, nil
	case Empty:
		return obj == nil, nil
	case NotEmpty:
		return obj != nil, nil
	case InRange:
		ranges := []string{}
		err := parseJSON(values, &ranges)
		if err != nil {
			return false, err
		}
		if len(ranges) != 2 {
			return false, fmt.Errorf("type/golang/date.go: values is invalid, %s", values)
		}
		from, err := time.Parse(time.RFC3339, ranges[0])
		if err != nil {
			return false, err
		}
		to, err := time.Parse(time.RFC3339, ranges[1])
		if err != nil {
			return false, err
		}
		return object.After(from) && object.Before(to), nil
	case NotInRange:
		ranges := []string{}
		err := parseJSON(values, &ranges)
		if err != nil {
			return false, err
		}
		if len(ranges) != 2 {
			return false, fmt.Errorf("type/golang/date.go: values is invalid, %s", values)
		}
		from, err := time.Parse(time.RFC3339, ranges[0])
		if err != nil {
			return false, err
		}
		to, err := time.Parse(time.RFC3339, ranges[1])
		if err != nil {
			return false, err
		}
		return object.Before(from) || object.After(to), nil
	case Before:
		value := ""
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		v, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return false, err
		}
		return object.Before(v), nil
	case After:
		value := ""
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		v, err := time.Parse(time.RFC3339, value)
		if err != nil {
			return false, err
		}
		return object.After(v), nil
	}

	return false, fmt.Errorf("type/golang/date.go: unsupport operator, %v, %s, %s", obj, operator, values)
}
