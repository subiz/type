package typesystem

import (
	"fmt"
	json "github.com/pquerna/ffjson/ffjson"
)

// StringsType set of strings type system
type StringsType struct {
	stringtype iType
}

// NewStringsType create new set of strings type system
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

func (t *StringsType) Evaluate(obj interface{} /*slice*/, operator string, values string) bool {
	var object []string
	if obj != nil {
		var ok bool
		object, ok = obj.([]string)
		if !ok {
			fmt.Printf("type/golang/strings.go: obj must be a slice\n")
			return false
		}
	}
	switch operator {
	case Empty:
		return len(object) == 0
	case NotEmpty:
		return len(object) != 0
	case Eq:
		value := make([]string, 0)
		json.Unmarshal([]byte(values), &value)
		return equals(object, value)
	case Ne:
		value := make([]string, 0)
		json.Unmarshal([]byte(values), &value)
		return equals(object, value)
	case Subset:
		value := make([]string, 0)
		json.Unmarshal([]byte(values), &value)
		return superset(value, object)
	case NotSubset:
		value := make([]string, 0)
		json.Unmarshal([]byte(values), &value)
		return !superset(value, object)
	case Superset:
		value := make([]string, 0)
		json.Unmarshal([]byte(values), &value)
		return superset(object, value)
	case NotSuperset:
		value := make([]string, 0)
		json.Unmarshal([]byte(values), &value)
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
		panic("unsupported operator: " + operator)
	}
}
