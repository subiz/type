package typesystem

import (
	"fmt"
	"regexp"
	"strings"
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

func (t *StringType) Evaluate(obj interface{}, operator string, values string) (bool, error) {
	var object string
	if obj != nil {
		var ok bool
		object, ok = obj.(string)
		if !ok {
			return false, fmt.Errorf("type/golang/string.go: obj must be a string, got `%v`\n", obj)
		}
	}
	switch operator {
	case Empty:
		return obj == nil || len(strings.TrimSpace(object)) == 0, nil
	case NotEmpty:
		return obj != nil && len(strings.TrimSpace(object)) != 0, nil
	case Eq:
		var value string
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return object == value, nil
	case Ne:
		var value string
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return object != value, nil
	case Regex:
		var value string
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return regexp.MatchString(value, object)
	case In:
		value := make([]string, 0)
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return contains(value, object), nil
	case NotIn:
		value := make([]string, 0)
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return !contains(value, object), nil
	case StartsWith:
		var value string
		fmt.Println("vao roi")
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		fmt.Println("obj-vl ", object, value)
		return strings.HasPrefix(object, value), nil
	case NotStartsWith:
		var value string
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return strings.HasPrefix(object, value), nil
	case EndsWith:
		var value string
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return strings.HasSuffix(object, value), nil
	case NotEndsWith:
		var value string
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return !strings.HasSuffix(object, value), nil
	case Contains:
		var value string
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return strings.Contains(object, value), nil
	case NotContains:
		var value string
		if err := parseJSON(values, &value); err != nil {
			return false, err
		}
		return !strings.Contains(object, value), nil
	default:
		return false, fmt.Errorf("type/golang/string.go: unsupport operator, %v, %s, %s", obj, operator, values)
	}
}

func (t *StringType) ConvToEls(key, operator, values string) (string, error) {
	var value string
	if err := parseJSON(values, &value); err != nil {
		return "", err
	}
	switch operator {
	case Eq:
		return fmt.Sprintf(`{"bool": {"must": {"match_phrase": { %q: %q}}}}`, key, value), nil
	case Ne:
		return fmt.Sprintf(`{"bool": {"must_not": {"match_phrase": { %q: %q}}}}`, key, value), nil
		// case In:
		// 	return fmt.Sprintf(`{"must": {"wildcard": { "%s", "*%s*"}}}`, key, value), nil
		// case NotIn:
		// return fmt.Sprintf(`{"must_not": {"wildcard": { "%s", "*%s*"}}}`, key, value), nil
	case StartsWith:
		return fmt.Sprintf(`{"bool": {"must": {"wildcard": { %q: %q}}}}`, key, value+"*"), nil
	case NotStartsWith:
		return fmt.Sprintf(`{"bool": {"must_not": {"wildcard": { %q: %q}}}}`, key, value+"*"), nil
	case EndsWith:
		return fmt.Sprintf(`{"bool": {"must": {"wildcard": { %q: %q}}}}`, key, "*"+value), nil
	case NotEndsWith:
		return fmt.Sprintf(`{"bool": {"must_not": {"wildcard": { %q: %q}}}}`, key, "*"+value), nil
	case Contains:
		return fmt.Sprintf(`{"bool": {"must": {"match": { %q: %q}}}}`, key, "*"+value+"*"), nil
	case NotContains:
		return fmt.Sprintf(`{"bool": {"must_not": {"wildcard": { %q: %q}}}}`, key, "*"+value+"*"), nil
	default:
		return "", fmt.Errorf("type/golang/string.go: unsupport operator, %v, %s, %s", key, operator, value)
	}
}
