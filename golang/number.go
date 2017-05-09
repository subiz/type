package typesystem

import (
	"strconv"
	"bitbucket.org/subiz/gocommon"
	"errors"
	"math"
)

const Tolerance = 0.000001

type NumberType struct {
}

func NewNumberType() iType {
	return &NumberType{}
}

func (t *NumberType) Evaluate(obj interface{}, operator string, values interface{}) bool {
	var pobj *string
	if obj != nil {
		var ok bool
		pobj, ok = obj.(*string)
		if !ok {
			common.Log("obj must be pointer to a string")
			return false
		}
	}
	var object float64
	var err error
	if pobj != nil {
		object, err = strconv.ParseFloat(*pobj, 64)
	}
	switch operator {
	case Nan:
		return err != nil
	case An:
		return err == nil
	case Empty:
		return pobj == nil
	case NotEmpty:
		return pobj != nil
	case Eq:
		if pobj == nil {
			return false
		}
		valuestring, ok := values.(string)
		if !ok {
			return false
		}
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return math.Abs(value - object) < Tolerance
	case Ne:
		if pobj == nil {
			return false
		}
		valuestring, ok := values.(string)
		if !ok {
			return false
		}
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return true
		}
		return math.Abs(value - object) > Tolerance
	case Gt:
		if pobj == nil {
			return false
		}
		valuestring, ok := values.(string)
		if !ok {
			return false
		}
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return value < object
	case Lt:
		if pobj == nil {
			return false
		}
		valuestring, ok := values.(string)
		if !ok {
			return false
		}
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return object < value
	case Gte:
		if pobj == nil {
			return false
		}
		valuestring, ok := values.(string)
		if !ok {
			return false
		}
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return value < object || math.Abs(value - object) < Tolerance
	case Lte:
		if pobj == nil {
			return false
		}
		valuestring, ok := values.(string)
		if !ok {
			return false
		}
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return object < value || math.Abs(value - object) < Tolerance
	case In:
		if pobj == nil {
			return false
		}
		valuesstring, ok := values.([]string)
		if !ok {
			return false
		}
		for _, s := range valuesstring {
			if  *pobj == s {
				return true
			}
		}
		return false
	case NotIn:
		if pobj == nil {
			return false
		}
		valuesstring, ok := values.([]string)
		if !ok {
			return false
		}
		for _, s := range valuesstring {
			if  *pobj == s {
				return false
			}
		}
		return true
	case InRange:
		if pobj == nil {
			return false
		}
		valuestring, ok := values.([]string)
		if !ok || len(valuestring) < 2 {
			return false
		}
		lower, err := strconv.ParseFloat(valuestring[0], 64)
		if err != nil {
			return false
		}
		upper, err := strconv.ParseFloat(valuestring[1], 64)
		if err != nil {
			return false
		}
		return lower < object && object < upper ||
			math.Abs(object - lower) < Tolerance ||
			math.Abs(object - upper) < Tolerance
	case NotInRange:
		if pobj == nil {
			return false
		}
		valuestring, ok := values.([]string)
		if !ok || len(valuestring) < 2 {
			return false
		}
		lower, err := strconv.ParseFloat(valuestring[0], 64)
		if err != nil {
			return false
		}
		upper, err := strconv.ParseFloat(valuestring[1], 64)
		if err != nil {
			return false
		}
		return object < lower || upper < object &&
			math.Abs(object - lower) > Tolerance &&
			math.Abs(object - upper) > Tolerance
	default:
		common.Panic(errors.New("unsupported operator"), "unsupported operator: " + operator)
	}
	return false
}
