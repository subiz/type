package typesystem

import (
	"fmt"
	"strconv"
)

// BoolType boolean type system
type BoolType struct {
}

// NewBoolType create new boolean type system
func NewBoolType() *BoolType {
	return &BoolType{}
}

func (t *BoolType) ConvToEls(key, operator, value string) (string, error) {
	switch operator {
	// case Nab:
	// 	return err != nil, nil
	// case Ab:
	// 	return err == nil, nil
	// case Empty:
	// 	return obj == nil, nil
	// case NotEmpty:
	// 	return obj != nil, nil
	case True:
		return fmt.Sprintf(`{"bool": {"must": {"term": { %q: %t}}}}`, key, true), nil
	case False:
		return fmt.Sprintf(`{"bool": {"must": {"term": { %q: %t}}}}`, key, false), nil
	default:
		return "", fmt.Errorf("type/golang/bool.go: unsupport operator, %v, %s, %s", key, operator, value)
	}
}

// Evaluate operator(object, values) == true
func (t *BoolType) Evaluate(obj interface{}, operator string, values string) (bool, error) {
	var pobj string
	var err error
	var object bool

	if obj != nil {
		if _, ok := obj.(bool); ok {
			if obj.(bool) {
				obj = "true"
			} else {
				obj = "false"
			}
		}
		var ok bool
		pobj, ok = obj.(string)
		if !ok {
			return false, fmt.Errorf("type/golang/bool.go: obj must be a string or a bool, got `%v`\n", obj)
		}

		object, err = strconv.ParseBool(pobj)
		if err != nil {
			return false, fmt.Errorf("type/golang/bool.go: %v unable to parse value `%s` to bool\n", err, pobj)
		}
	}
	switch operator {
	case Nab:
		return err != nil, nil
	case Ab:
		return err == nil, nil
	case Empty:
		return obj == nil, nil
	case NotEmpty:
		return obj != nil, nil
	case True:
		if obj == nil {
			return false, nil
		}
		return object, nil
	case False:
		if obj == nil {
			return true, nil
		}
		return !object, nil
	default:
		return false, fmt.Errorf("type/golang/bool.go: unsupport operator, %v, %s, %s", obj, operator, values)
	}
}
