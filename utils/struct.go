package utils

//TODO make a public utils package and put this in there so that i dont copy+paste this all the time...

import (
	"fmt"
	"reflect"
	"strings"
)

func StructKeys(obj any) (keys []string) {
	var val = reflect.Indirect(reflect.ValueOf(obj))
	var length = val.Type().NumField()
	keys = make([]string, length)

	for i := range length {
		keys[i] = val.Type().Field(i).Name
	}

	return keys
}
func StructValues(obj any) (vals []any) {
	var val = reflect.Indirect(reflect.ValueOf(obj))
	var length = val.Type().NumField()
	vals = make([]any, length)

	for i := range length {
		vals[i] = val.Field(i)
	}

	return vals
}

func StructToString(obj any) string {
	var lines []string
	var values = StructValues(obj)
	for i, key := range StructKeys(obj) {
		lines = append(lines, fmt.Sprintf("%s: %v", key, values[i]))
	}

	return strings.Join(lines, "\n")
}
