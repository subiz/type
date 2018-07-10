package typesystem

import (
	"strconv"
	json "github.com/pquerna/ffjson/ffjson"
	"math"
	"fmt"
)

const Tolerance = 0.000001

type NumberType struct {
}

func NewNumberType() iType {
	return &NumberType{}
}

func parseJSON(j string, out interface{}) {
	if err := json.Unmarshal([]byte(j), out); err != nil {
		panic(err)
	}
}

func (t *NumberType) Evaluate(obj interface{}, operator string, values string) bool {
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
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return false
		}
		return math.Abs(value - object) < Tolerance
	case Ne:
		if obj == nil {
			return false
		}
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return true
		}
		return math.Abs(value - object) > Tolerance
	case Gt:
		if obj == nil {
			return false
		}
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return false
		}
		return value < object
	case Lt:
		if obj == nil {
			return false
		}
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return false
		}
		return object < value
	case Gte:
		if obj == nil {
			return false
		}
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return false
		}
		return value < object || math.Abs(value - object) < Tolerance
	case Lte:
		if obj == nil {
			return false
		}
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return false
		}
		return object < value || math.Abs(value - object) < Tolerance
	case In:
		if obj == nil {
			return false
		}
		fs := make([]float64, 0)
		parseJSON(values, &fs)
		for _, f := range fs {
			if math.Abs(f - object) < Tolerance {
				return true
			}
		}
		return false
	case NotIn:
		if obj == nil {
			return false
		}
		fs := make([]float64, 0)
		parseJSON(values, &fs)
		for _, f := range fs {
			if math.Abs(f - object) < Tolerance {
				return false
			}
		}
		return true
	case InRange:
		if obj == nil {
			return false
		}
		fs := make([]float64, 0)
		parseJSON(values, &fs)
		if len(fs) < 2 {
			return false
		}
		lower, upper := fs[0], fs[1]
		return lower < object && object < upper ||
			math.Abs(object - lower) < Tolerance ||
			math.Abs(object - upper) < Tolerance
	case NotInRange:
		if obj == nil {
			return false
		}
		fs := make([]float64, 0)
		parseJSON(values, &fs)
		vs := convertToInterfaceSlice(values)
		if len(vs) < 2 {
			return false
		}
		lower, upper := fs[0], fs[1]
		return object < lower || upper < object &&
			math.Abs(object - lower) > Tolerance &&
			math.Abs(object - upper) > Tolerance
	default:
		panic("unsupported operator: " + operator)
	}
}
