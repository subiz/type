package typesystem

import (
	"strconv"
	"fmt"
)

// BoolType boolean type system
type BoolType struct {
}

// NewBoolType create new boolean type system
func NewBoolType() *BoolType {
	return &BoolType{}
}

// Evaluate operator(object, values) == true
func (t *BoolType) Evaluate(obj interface{}, operator string, values string) bool {
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
			fmt.Printf("type/golang/bool.go: obj must be a string or a bool, got `%v`\n", obj)
			return false
		}

		object, err = strconv.ParseBool(pobj)
		if err != nil {
			fmt.Printf("type/golang/bool.go: %v unable to parse value `%s` to bool\n", err, pobj)
			return false
		}
	}
	switch operator {
	case Nab:
		return err != nil
	case Ab:
		return err == nil
	case Empty:
		return obj == nil
	case NotEmpty:
		return obj != nil
	case True:
		if obj == nil {
			return false
		}
		return object
	case False:
		if obj == nil {
			return true
		}
		return !object
	default:
		panic("unsupported operator: " + operator)
	}
}
