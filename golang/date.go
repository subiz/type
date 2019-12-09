package typesystem

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type DateType struct {
}

func NewDateType() iType {
	return &DateType{}
}

func (me *DateType) ConvToEls(key, operator, values string) (string, error) {
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
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"gt": %q, "format": "date_time||date_time_no_millis"}}}}}`, key, ds), nil
	case Lt:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"lt": %q, "format": "date_time||date_time_no_millis"}}}}}`, key, ds), nil
	case Gte:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"gte": %q, "format": "date_time||date_time_no_millis"}}}}}`, key, ds), nil
	case Lte:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"lte": %q, "format": "date_time||date_time_no_millis"}}}}}`, key, ds), nil
	case InRange:
		fs := make([]string, 0)
		if err := parseJSON(values, &fs); err != nil {
			return "", err
		}
		if len(fs) < 2 {
			return "", errors.New("Worng format")
		}
		lower, upper := fs[0], fs[1]
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"gte": %q, "lte": %q, "format": "date_time||date_time_no_millis"}}}}}`, key, lower, upper), nil
	case NotInRange:
		fs := make([]string, 0)
		if err := parseJSON(values, &fs); err != nil {
			return "", err
		}
		if len(fs) < 2 {
			return "", errors.New("Wrong format")
		}
		lower, upper := fs[0], fs[1]
		return fmt.Sprintf(`{"bool": {"must_not": {"range": {%q: {"gte": %q, "lte": %q, "format": "date_time||date_time_no_millis"}}}}}`, key, lower, upper), nil
	case After:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"gt": %q, "format": "date_time||date_time_no_millis"}}}}}`, key, ds), nil
	case Before:
		var ds string
		if err := parseJSON(values, &ds); err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"bool": {"must": {"range": {%q: {"lt": %q, "format": "date_time||date_time_no_millis"}}}}}`, key, ds), nil
	default:
		return "", fmt.Errorf("type/golang/datetime.go: unsupport operator, %v, %s, %s", key, operator, values)
	}
}

func parseTime(jsonEncoded string) (time.Time, error) {
	var t time.Time
	var s string
	err := json.Unmarshal([]byte(jsonEncoded), &s)
	if err != nil {
		return t, err
	}
	t, err = time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return t, err
	}
	return t, nil
}

func parseTimeRange(jsonEncoded string) ([2]time.Time, error) {
	rang := [2]time.Time{}
	ss := make([]string, 0)
	err := json.Unmarshal([]byte(jsonEncoded), &ss)
	if err != nil {
		return rang, err
	}
	if len(ss) < 2 {
		return rang, errors.New("Wrong format")
	}
	rang[0], err = time.Parse(time.RFC3339Nano, ss[0])
	if err != nil {
		return rang, err
	}
	rang[1], err = time.Parse(time.RFC3339Nano, ss[1])
	if err != nil {
		return rang, err
	}
	return rang, nil
}

func (me *DateType) ToBigQuery(key, operator, values string) (string, error) {
	switch operator {
	case Eq:
		t, err := parseTime(values)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`(%s = %d)`, key, t.UnixNano()/1e6), nil
	case Ne:
		t, err := parseTime(values)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`(%s != %d)`, key, t.UnixNano()/1e6), nil
	case Gt, After:
		t, err := parseTime(values)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`(%s > %d)`, key, t.UnixNano()/1e6), nil
	case Lt, Before:
		t, err := parseTime(values)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`(%s < %d)`, key, t.UnixNano()/1e6), nil
	case Gte:
		t, err := parseTime(values)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`(%s >= %d)`, key, t.UnixNano()/1e6), nil
	case Lte:
		t, err := parseTime(values)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`(%s <= %d)`, key, t.UnixNano()/1e6), nil
	case InRange:
		rang, err := parseTimeRange(values)
		if err != nil {
			return "", err
		}
		from, to := rang[0].UnixNano()/1e6, rang[1].UnixNano()/1e6
		return fmt.Sprintf(`(%s > %d AND %s < %d)`, key, from, key, to), nil
	case NotInRange:
		rang, err := parseTimeRange(values)
		if err != nil {
			return "", err
		}
		from, to := rang[0].UnixNano()/1e6, rang[1].UnixNano()/1e6
		return fmt.Sprintf(`(%s < %d OR %s > %d)`, key, from, key, to), nil
	default:
		return "", fmt.Errorf("type/golang/datetime.go: unsupport operator, %v, %s, %s", key, operator, values)
	}
}

// values is in json format
func (me *DateType) Evaluate(obj interface{}, operator string, values string) (bool, error) {
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
		rang, err := parseTimeRange(values)
		if err != nil {
			return false, err
		}
		return object.After(rang[0]) && object.Before(rang[1]), nil
	case NotInRange:
		rang, err := parseTimeRange(values)
		if err != nil {
			return false, err
		}
		return object.Before(rang[0]) || object.After(rang[1]), nil
	case Lt, Before:
		t, err := parseTime(values)
		if err != nil {
			return false, err
		}
		return object.Before(t), nil
	case Gt, After:
		t, err := parseTime(values)
		if err != nil {
			return false, err
		}
		return object.After(t), nil
	}

	return false, fmt.Errorf("type/golang/date.go: unsupport operator, %v, %s, %s", obj, operator, values)
}
