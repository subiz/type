package typesystem

import (
	"fmt"
	json "github.com/pquerna/ffjson/ffjson"
	"reflect"
)

func convertToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func convertToStringSlice(slice interface{}) []string {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil
	}
	ret := make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = fmt.Sprintf("%v", s.Index(i).Interface())
	}
	return ret
}

func parseJSON(j string, out interface{}) error {
	if err := json.Unmarshal([]byte(j), out); err != nil {
		return err
	}
	return nil
}
