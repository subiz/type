package typesystem

import (
	"strconv"
	"bitbucket.org/subiz/gocommon"
	"errors"
	"math"
	"fmt"
	"reflect"
)

const Tolerance = 0.000001

type NumberType struct {
}

func NewNumberType() iType {
	return &NumberType{}
}

func (t *NumberType) Evaluate(obj interface{}, operator string, values interface{}) bool {
	sobj := fmt.Sprintf("%v", obj)
	var object float64
	var err error
	if obj != nil {
		object, err = strconv.ParseFloat(sobj, 64)
	}
	switch operator {
	case Nan:
		return err != nil
	case An:
		return err == nil
	case Empty:
		return obj == nil
	case NotEmpty:
		return obj != nil
	case Eq:
		if obj == nil {
			return false
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return math.Abs(value - object) < Tolerance
	case Ne:
		if obj == nil {
			return false
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return true
		}
		return math.Abs(value - object) > Tolerance
	case Gt:
		if obj == nil {
			return false
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return value < object
	case Lt:
		if obj == nil {
			return false
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return object < value
	case Gte:
		if obj == nil {
			return false
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return value < object || math.Abs(value - object) < Tolerance
	case Lte:
		if obj == nil {
			return false
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false
		}
		return object < value || math.Abs(value - object) < Tolerance
	case In:
		if obj == nil {
			return false
		}
		vs := convertToInterfaceSlice(values)
		if vs == nil {
			return false
		}
		for _, s := range vs {
			s := fmt.Sprintf("%v", s)
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				continue
			}
			if math.Abs(v - object) < Tolerance {
				return true
			}
		}
		return false
	case NotIn:
		if obj == nil {
			return false
		}
		vs := convertToInterfaceSlice(values)
		if vs == nil {
			return false
		}
		for _, s := range vs {
			s := fmt.Sprintf("%v", s)
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				continue
			}
			if math.Abs(v - object) < Tolerance {
				return false
			}
		}
		return true
	case InRange:
		if obj == nil {
			return false
		}
		vs := convertToInterfaceSlice(values)
		if vs == nil || len(vs) < 2 {
			return false
		}
		lower, err := strconv.ParseFloat(fmt.Sprintf("%v", vs[0]), 64)
		if err != nil {
			return false
		}
		upper, err := strconv.ParseFloat(fmt.Sprintf("%v", vs[1]), 64)
		if err != nil {
			return false
		}
		return lower < object && object < upper ||
			math.Abs(object - lower) < Tolerance ||
			math.Abs(object - upper) < Tolerance
	case NotInRange:
		if obj == nil {
			return false
		}
		vs := convertToInterfaceSlice(values)
		if vs == nil || len(vs) < 2 {
			return false
		}
		lower, err := strconv.ParseFloat(fmt.Sprintf("%v", vs[0]), 64)
		if err != nil {
			return false
		}
		upper, err := strconv.ParseFloat(fmt.Sprintf("%v", vs[1]), 64)
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


func convertToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}
	ret := make([]interface{}, s.Len())
	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}
