package typesystem

import (
	"strings"
	"regexp"
	"bitbucket.org/subiz/gocommon"
)

// StringType string type system
type StringType struct {
}

// NewStringType create new string type
func NewStringType() *StringType {
	return &StringType{}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.ToLower(strings.Trim(a, " ")) == strings.ToLower(strings.Trim(e, " ")) {
			return true
		}
	}
	return false
}

func (t *StringType) Evaluate(obj interface{}, operator string, values interface{}) bool {
	var object string
	if obj != nil {
		var ok bool
		object, ok = obj.(string)
		if !ok {
			common.Logf("obj must be a string, got `%v`", obj)
			return false
		}
	}
	switch operator {
	case Empty:
		return obj == nil || len(strings.Trim(object, " ")) == 0
	case NotEmpty:
		return obj != nil && len(strings.Trim(object, " ")) != 0
	case Eq:
		value, ok := values.(string)
		return ok && object == value
	case Ne:
		value, ok := values.(string)
		return !ok || object != value
	case Regex:
		value, ok := values.(string)
		if !ok || obj == nil {
			return false
		}
		matched, err := regexp.MatchString(value, object)
		if err != nil {
			panic("failed regex")
		}
		return matched
	case In:
		value := convertToStringSlice(values)
		if value == nil {
			return false
		}
		return contains(value, object)
	case NotIn:
		value := convertToStringSlice(values)
		if value == nil {
			return true
		}
		return !contains(value, object)
	case StartsWith:
		value, ok := values.(string)
		if !ok || obj == nil {
			return false
		}
		return  strings.HasPrefix(object, value)
	case NotStartsWith:
		value, ok := values.(string)
		if !ok || obj == nil {
			return false
		}
		return  strings.HasPrefix(object, value)
	case EndsWith:
		value, ok := values.(string)
		if !ok || obj == nil {
			return false
		}
		return strings.HasSuffix(object, value)
	case NotEndsWith:
		value, ok := values.(string)
		if !ok || obj == nil {
			return false
		}
		return !strings.HasSuffix(object, value)
	case Contains:
		value, ok := values.(string)
		if !ok || obj == nil {
			return false
		}
		return strings.Contains(value, object)
	case NotContains:
		value, ok := values.(string)
		if !ok || obj == nil {
			return false
		}
		return !strings.Contains(value, object)
	default:
		panic("unsupported operator: " + operator)
	}
}
