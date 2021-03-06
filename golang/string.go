package typesystem

import (
	"encoding/json"
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

func (me *StringType) Evaluate(obj interface{}, operator string, values string) (bool, error) {
	if values == "" {
		values = `""`
	}
	var object string
	if obj != nil {
		var ok bool
		object, ok = obj.(string)
		if !ok {
			return false, fmt.Errorf("type/golang/string.go: obj must be a string, got `%v`", obj)
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

func (me *StringType) ConvToEls(key, operator, values string) (string, error) {
	if values == "" {
		values = `""`
	}

	var value string
	if err := parseJSON(values, &value); err != nil {
		return "", err
	}

	switch operator {
	case Eq:
		return fmt.Sprintf(`{"bool": {"must": {"match_phrase": {%q: %q}}}}`, key, value), nil
	case Ne:
		return fmt.Sprintf(`{"bool": {"must_not": {"match_phrase": {%q: %q}}}}`, key, value), nil
	case Contains:
		return fmt.Sprintf(`{"bool": {"must": {"match": { %q: %q }}}}`, key, value), nil
	case NotContains:
		return fmt.Sprintf(`{"bool": {"must_not": {"match": {%q: %q }}}}`, key, value), nil
	case Empty:
		return fmt.Sprintf(`{"bool": {"must_not": {"exists": {"field": %q}}}}`, key), nil
	case NotEmpty:
		return fmt.Sprintf(`{"bool": {"must": {"exists": {"field": %q}}}}`, key), nil
	default:
		return "", fmt.Errorf("type/golang/string.go: unsupport operator, %v, %s, %s", key, operator, value)
	}
}

func (me *StringType) ToBigQuery(key, operator, values string) (string, error) {
	if values == "" {
		values = `""`
	}

	var value string
	if err := json.Unmarshal([]byte(values), &value); err != nil {
		return "", err
	}

	switch operator {
	case Eq:
		return fmt.Sprintf(`(%s = %q)`, key, value), nil
	case Ne:
		return fmt.Sprintf(`(%s != %q)`, key, value), nil
	case Contains:
		return fmt.Sprintf(`(%s LIKE "%s")`, key, value), nil
	case NotContains:
		return fmt.Sprintf(`(%s NOT LIKE "%s")`, key, value), nil
	case Empty:
		return fmt.Sprintf(`(%s IS NULL OR %s = "")`, key, key), nil
	case NotEmpty:
		return fmt.Sprintf(`(%s IS NOT NULL AND %s != "")`, key, key), nil
	default:
		return "", fmt.Errorf("type/golang/string.go: unsupport operator, %v, %s, %s", key, operator, value)
	}
}
