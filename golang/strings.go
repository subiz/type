package typesystem

import (
	"fmt"
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

func (t *StringsType) ConvToEls(key, operator, value string) (string, error) {
	switch operator {
	case Contains:
		return fmt.Sprintf(`{"bool":{"must": {"term": { %q: {"value": %q }}}}}`, key, value), nil
	case NotContains:
		return fmt.Sprintf(`{"bool":{"must_not": {"term": {%q: {"value": %q }}}}}`, key, value), nil
	default:
		return "", fmt.Errorf("type/golang/strings.go: unsupport operator, %v, %s, %s", key, operator, value)
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

func (t *StringsType) Evaluate(obj interface{} /*slice*/, operator string, values string) (bool, error) {
	var object []string
	if obj != nil {
		var ok bool
		object, ok = obj.([]string)
		if !ok {
			return false, fmt.Errorf("type/golang/strings.go: obj must be a slice\n")
		}
	}
	switch operator {
	case Empty:
		return len(object) == 0, nil
	case NotEmpty:
		return len(object) != 0, nil
	case Eq:
		value := make([]string, 0)
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return equals(object, value), nil
	case Ne:
		value := make([]string, 0)
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return equals(object, value), nil
	case Subset:
		value := make([]string, 0)
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return superset(value, object), nil
	case NotSubset:
		value := make([]string, 0)
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return !superset(value, object), nil
	case Superset:
		value := make([]string, 0)
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return superset(object, value), nil
	case NotSuperset:
		value := make([]string, 0)
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return !superset(object, value), nil
	case Regex:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, Regex, values); o {
				return true, nil
			}
		}
		return false, nil
	case In:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, In, values); o {
				return true, nil
			}
		}
		return false, nil
	case NotIn:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, In, values); o {
				return false, nil
			}
		}
		return true, nil
	case StartsWith:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, StartsWith, values); o {
				return true, nil
			}
		}
		return false, nil
	case EndsWith:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, EndsWith, values); o {
				return true, nil
			}
		}
		return false, nil
	case NotStartsWith:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, StartsWith, values); o {
				return false, nil
			}
		}
		return true, nil
	case NotEndsWith:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, EndsWith, values); o {
				return false, nil
			}
		}
		return true, nil
	case Contains:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, Contains, values); o {
				return true, nil
			}
		}
		return false, nil
	case NotContains:
		for _, s := range object {
			if o, _ := t.stringtype.Evaluate(s, Contains, values); o {
				return false, nil
			}
		}
		return true, nil
	default:
		return false, fmt.Errorf("type/golang/strings.go: unsupport operator, %v, %s, %s", obj, operator, values)
	}
}
