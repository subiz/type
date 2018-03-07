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
		if strings.ToLower(strings.TrimSpace(a)) == strings.ToLower(strings.TrimSpace(e)) {
			return true
		}
	}
	return false
}

func (t *StringType) Evaluate(obj interface{}, operator string, values string) bool {
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
		return obj == nil || len(strings.TrimSpace(object)) == 0
	case NotEmpty:
		return obj != nil && len(strings.TrimSpace(object)) != 0
	case Eq:
		var value string
		common.ParseJSON(values, &value)
		return object == value
	case Ne:
		var value string
		common.ParseJSON(values, &value)
		return object != value
	case Regex:
		var value string
		common.ParseJSON(values, &value)
		matched, err := regexp.MatchString(value, object)
		if err != nil {
			panic("failed regex")
		}
		return matched
	case In:
		value := make([]string, 0)
		common.ParseJSON(values, &value)
		return contains(value, object)
	case NotIn:
		value := make([]string, 0)
		common.ParseJSON(values, &value)
		return !contains(value, object)
	case StartsWith:
		var value string
		common.ParseJSON(values, &value)
		return  strings.HasPrefix(object, value)
	case NotStartsWith:
		var value string
		common.ParseJSON(values, &value)
		return  strings.HasPrefix(object, value)
	case EndsWith:
		var value string
		common.ParseJSON(values, &value)
		return strings.HasSuffix(object, value)
	case NotEndsWith:
		var value string
		common.ParseJSON(values, &value)
		return !strings.HasSuffix(object, value)
	case Contains:
		var value string
		return strings.Contains(object, value)
	case NotContains:
		var value string
		common.ParseJSON(values, &value)
		return !strings.Contains(object, value)
	default:
		panic("unsupported operator: " + operator)
	}
}
