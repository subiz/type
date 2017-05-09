package typesystem

import (
	"strings"
	"regexp"
	"bitbucket.org/subiz/gocommon"
	"errors"
)

type StringType struct {
}

func NewStringType() iType {
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
	var pobj *string
	if obj != nil {
		var ok bool
		pobj, ok = obj.(*string)
		if !ok {
			common.Log("obj must be pointer to a string")
			return false
		}
	}

	if pobj != nil {
		object = *pobj
	}

	switch operator {
	case Empty:
		return pobj == nil || len(strings.Trim(object, " ")) == 0
	case NotEmpty:
		return pobj != nil && len(strings.Trim(object, " ")) != 0
	case Eq:
		value, ok := values.(string)
		return ok && object == value
	case Ne:
		value, ok := values.(string)
		return !ok || object != value
	case Regex:
		value, ok := values.(string)
		if !ok || pobj == nil {
			return false
		}
		matched, err := regexp.MatchString(value, object)
		common.Panic(err, "failed regex")
		return matched
	case In:
		value, ok := values.([]string)
		if !ok || pobj == nil {
			return false
		}
		return contains(value, object)
	case NotIn:
		value, ok := values.([]string)
		if !ok || pobj == nil {
			return false
		}
		return !contains(value, object)
	case StartsWith:
		value, ok := values.(string)
		if !ok || pobj == nil {
			return false
		}
		return  strings.HasPrefix(object, value)
	case NotStartsWith:
		value, ok := values.(string)
		if !ok || pobj == nil {
			return false
		}
		return  strings.HasPrefix(object, value)
	case EndsWith:
		value, ok := values.(string)
		if !ok || pobj == nil {
			return false
		}
		return strings.HasSuffix(object, value)
	case NotEndsWith:
		value, ok := values.(string)
		if !ok || pobj == nil {
			return false
		}
		return !strings.HasSuffix(object, value)
	case Contains:
		value, ok := values.(string)
		if !ok || pobj == nil {
			return false
		}
		return strings.Contains(value, object)
	case NotContains:
		value, ok := values.(string)
		if !ok || pobj == nil {
			return false
		}
		return !strings.Contains(value, object)
	default:
		common.Panic(errors.New("unsupported operator"), "unsupported operator: " + operator)
	}
	return false
}
