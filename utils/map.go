package utils

import (
	"fmt"
	"strings"
)

func MapToString[K comparable, V any](obj map[K]V, stringConvert func(val V) string) string {
	if len(obj) == 0 {
		return fmt.Sprintf("%T { }", obj)
	}

	var lines []string
	for key, item := range obj {
		var itemLines = strings.Split(stringConvert(item), "\n")
		lines = append(lines, fmt.Sprintf("%v: {\n  %s\n }", key, strings.Join(itemLines, "\n  ")))
	}

	return fmt.Sprintf("%T {\n %s\n}", obj, strings.Join(lines, "\n "))
}
