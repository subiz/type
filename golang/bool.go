package typesystem

import (
	"strconv"
	"bitbucket.org/subiz/gocommon"
)

// BoolType boolean type system
type BoolType struct {
}

// NewBoolType create new boolean type system
func NewBoolType() *BoolType {
	return &BoolType{}
}

// Evaluate operator(object, values) == true
func (t *BoolType) Evaluate(obj interface{}, operator string, values interface{}) bool {
	var pobj string
	if obj != nil {
		var ok bool
		pobj, ok = obj.(string)
		if !ok {
			common.Logf("obj must be a string, got `%s`", obj)
			return false
		}
	}
	var err error
	var object bool
	if obj != nil {
		object, err = strconv.ParseBool(pobj)
		if err != nil {
			common.Logf("%v unable to parse value `%s` to bool", err, pobj)
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
			return false
		}
		return !object
	default:
		panic("unsupported operator: " + operator)
	}
}
