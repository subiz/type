package typesystem

import (
	"bitbucket.org/subiz/gocommon"
	"errors"
)

type StringsType struct {
	stringtype iType
}

func NewStringsType() iType {
	return &StringsType{
		stringtype: NewStringType(),
	}
}

func superset(a []string, b []string) bool {
	for _, i := range b {
		found := false
		for _, j := range a {
			if i == j {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func equals(a []string, b []string) bool {
	return superset(a, b) && superset(b, a)
}

func (t *StringsType) Evaluate(obj interface{} /*slice*/, operator string, values interface{}) bool {
	var object []string
	if obj != nil {
		var ok bool
		object, ok = obj.([]string)
		if !ok {
			common.Log("obj must be a slice")
			return false
		}
	}
	switch operator {
	case Empty:
		return len(object) == 0
	case NotEmpty:
		return len(object) != 0
	case Eq:
		value, ok := values.([]string)
		if !ok {
			return false
		}
		return equals(object, value)
	case Ne:
		value, ok := values.([]string)
		if !ok {
			return true
		}
		return equals(object, value)
	case Subset:
		value, ok := values.([]string)
		if !ok {
			return false
		}
		return superset(value, object)
	case NotSubset:
		value, ok := values.([]string)
		if !ok {
			return true
		}
		return !superset(value, object)
	case Superset:
		value, ok := values.([]string)
		if !ok {
			return false
		}
		return superset(object, value)
	case NotSuperset:
		value, ok := values.([]string)
		if !ok {
			return true
		}
		return !superset(object, value)
	case Regex:
		for _, s := range object {
			if t.stringtype.Evaluate(s, Regex, values) {
				return true
			}
		}
		return false
	case In:
		for _, s := range object {
			if t.stringtype.Evaluate(s, In, values) {
				return true
			}
		}
		return false
	case NotIn:
		for _, s := range object {
			if t.stringtype.Evaluate(s, In, values) {
				return false
			}
		}
		return true
	case StartsWith:
		for _, s := range object {
			if t.stringtype.Evaluate(s, StartsWith, values) {
				return true
			}
		}
		return false
	case EndsWith:
		for _, s := range object {
			if t.stringtype.Evaluate(s, EndsWith, values) {
				return true
			}
		}
		return false
	case NotStartsWith:
		for _, s := range object {
			if t.stringtype.Evaluate(s, StartsWith, values) {
				return false
			}
		}
		return true
	case NotEndsWith:
		for _, s := range object {
			if t.stringtype.Evaluate(s, EndsWith, values) {
				return false
			}
		}
		return true
	case Contains:
		for _, s := range object {
			if t.stringtype.Evaluate(s, Contains, values) {
				return true
			}
		}
		return false
	case NotContains:
		for _, s := range object {
			if t.stringtype.Evaluate(s, Contains, values) {
				return false
			}
		}
		return true
	default:
		common.Panic(errors.New("unsupported operator"), "unsupported operator: " + operator)
	}
	return false
}
