package typesystem

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type DateType struct {
}

func NewDateType() iType {
	return &DateType{}
}

func (t *DateType) ConvToEls(key, operator, values string) (string, error) {
	switch operator {
	// case Nan:
	// 	return err != nil, nil
	// case An:
	// 	return err == nil, nil
	// case Empty:
	// 	return obj == nil, nil
	// case NotEmpty:
	// 	return obj != nil, nil
	case Eq:
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"must": {"term": { "%s", %f}}}`, key, value), nil
	case Ne:
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"must_not": {"term": {"%s", %f}}}`, key, value), nil
	case Gt:
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"must": {"range": {"%s": {"gt": %f}}}}`, key, value), nil
	case Lt:
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"must": {"range": {"%s": {"lt": %f}}}}`, key, value), nil
	case Gte:
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"must": {"range": {"%s": {"gte": %f}}}}`, key, value), nil
	case Lte:
		value, err := strconv.ParseFloat(values, 64)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf(`{"must": {"range": {"%s": {"lte": %f}}}}`, key, value), nil
	case InRange:
		fs := make([]float64, 0)
		if err := parseJSON(values, &fs); err != nil {
			return "", err
		}
		if len(fs) < 2 {
			return "", errors.New("Worng format")
		}
		lower, upper := fs[0], fs[1]
		return fmt.Sprintf(`{"must": {"range": {"%s" {"gte": %f, "lte": %f}}}}`, key, lower, upper), nil
	case NotInRange:
		fs := make([]float64, 0)
		if err := parseJSON(values, &fs); err != nil {
			return "", err
		}
		if len(fs) < 2 {
			return "", errors.New("Worng format")
		}
		lower, upper := fs[0], fs[1]
		return fmt.Sprintf(`{"must_not": {"range": {"%s" {"gte": %f, "lte": %f}}}}`, key, lower, upper), nil
	default:
		return "", fmt.Errorf("type/golang/number.go: unsupport operator, %v, %s, %s", key, operator, values)
	}
}

// values is in json format
func (t *DateType) Evaluate(obj interface{}, operator string, values string) (bool, error) {
	sobj := fmt.Sprintf("%v", obj)
	var object float64
	var err error
	if obj != nil {
		object, err = strconv.ParseFloat(sobj, 64)
	}
	switch operator {
	case Nad:
		return err != nil, nil
	case An:
		return err == nil, nil
	case Empty:
		return obj == nil, nil
	case NotEmpty:
		return obj != nil, nil
	case Eq:
		if obj == nil {
			return false, nil
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false, nil
		}
		return math.Abs(value-object) < Tolerance, nil
	case Ne:
		if obj == nil {
			return false, nil
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return true, nil
		}
		return math.Abs(value-object) > Tolerance, nil
	case Gt:
		if obj == nil {
			return false, nil
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false, nil
		}
		return value < object, nil
	case Lt:
		if obj == nil {
			return false, nil
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false, nil
		}
		return object < value, nil
	case Gte:
		if obj == nil {
			return false, nil
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false, nil
		}
		return value < object || math.Abs(value-object) < Tolerance, nil
	case Lte:
		if obj == nil {
			return false, nil
		}
		valuestring := fmt.Sprintf("%v", values)
		value, err := strconv.ParseFloat(valuestring, 64)
		if err != nil {
			return false, nil
		}
		return object < value || math.Abs(value-object) < Tolerance, nil
	case In:
		if obj == nil {
			return false, nil
		}
		vs := convertToInterfaceSlice(values)
		if vs == nil {
			return false, nil
		}
		for _, s := range vs {
			s := fmt.Sprintf("%v", s)
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				continue
			}
			if math.Abs(v-object) < Tolerance {
				return true, nil
			}
		}
		return false, nil
	case NotIn:
		if obj == nil {
			return false, nil
		}
		vs := convertToInterfaceSlice(values)
		if vs == nil {
			return false, nil
		}
		for _, s := range vs {
			s := fmt.Sprintf("%v", s)
			v, err := strconv.ParseFloat(s, 64)
			if err != nil {
				continue
			}
			if math.Abs(v-object) < Tolerance {
				return false, nil
			}
		}
		return true, nil
	case InRange:
		if obj == nil {
			return false, nil
		}
		vs := convertToInterfaceSlice(values)
		if vs == nil || len(vs) < 2 {
			return false, nil
		}
		lower, err := strconv.ParseFloat(fmt.Sprintf("%v", vs[0]), 64)
		if err != nil {
			return false, nil
		}
		upper, err := strconv.ParseFloat(fmt.Sprintf("%v", vs[1]), 64)
		if err != nil {
			return false, nil
		}
		return lower < object && object < upper ||
			math.Abs(object-lower) < Tolerance ||
			math.Abs(object-upper) < Tolerance, nil
	case NotInRange:
		if obj == nil {
			return false, nil
		}
		vs := convertToInterfaceSlice(values)
		if vs == nil || len(vs) < 2 {
			return false, nil
		}
		lower, err := strconv.ParseFloat(fmt.Sprintf("%v", vs[0]), 64)
		if err != nil {
			return false, nil
		}
		upper, err := strconv.ParseFloat(fmt.Sprintf("%v", vs[1]), 64)
		if err != nil {
			return false, nil
		}
		return object < lower || upper < object &&
			math.Abs(object-lower) > Tolerance &&
			math.Abs(object-upper) > Tolerance, nil
	default:
		return false, fmt.Errorf("type/golang/date.go: unsupport operator, %v, %s, %s", obj, operator, values)
	}
}
