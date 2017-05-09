package typesystem

import (
	"strconv"
	"bitbucket.org/subiz/gocommon"
	"errors"
)

type BoolType struct {
}

func NewBoolType() iType {
	return &BoolType{}
}


func (t *BoolType) Evaluate(obj interface{}, operator string, values interface{}) bool {
	var pobj *string
	if obj != nil {
		var ok bool
		pobj, ok = obj.(*string)
		if !ok {
			common.Log("obj must be pointer to a string")
			return false
		}
	}
	var err error
	var object bool
	if pobj != nil {
		object, err = strconv.ParseBool(*pobj)
	}
	switch operator {
	case Nab:
		return err != nil
	case Ab:
		return err == nil
	case Empty:
		return pobj == nil
	case NotEmpty:
		return pobj != nil
	case True:
		if pobj == nil {
			return false
		}
		return object
	case False:
		if pobj == nil {
			return false
		}
		return !object
	default:
		common.Panic(errors.New("unsupported operator"), "unsupported operator: " + operator)
	}
	return false
}
