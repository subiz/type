package typesystem

import (
	"fmt"
	"time"
)

type DateType struct {
}

func NewDateType() iType {
	return &DateType{}
}

func (t *DateType) ConvToEls(key, operator, values string) (string, error) {
	return "", fmt.Errorf("type/golang/date.go: unsupport operator, %v, %s, %s", key, operator, values)
}

// values is in json format
func (t *DateType) Evaluate(obj interface{}, operator string, values string) (bool, error) {
	sobj := fmt.Sprintf("%v", obj)
	var object time.Time
	var err error
	if obj != nil {
		object, err = time.Parse(time.RFC3339, sobj)
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
		v, err := time.Parse(time.RFC3339, values)
		if err != nil {
			return false, err
		}
		return object.Before(v), nil
	case After:
		v, err := time.Parse(time.RFC3339, values)
		if err != nil {
			return false, err
		}
		return object.After(v), nil
	}

	return false, fmt.Errorf("type/golang/date.go: unsupport operator, %v, %s, %s", obj, operator, values)
}
