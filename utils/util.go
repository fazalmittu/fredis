package utils

import (
	"fmt"
)

func FormatValue(v interface{}) string {
	// switch reflect.TypeOf(v).Kind() {
	// case reflect.String:
	// 	s, ok := v.(string)
	// 	if ok {
	// 		return s
	// 	}
	// case reflect.Int:

	// }
	switch val := v.(type) {
	case string:
		return val
	case int:
		return fmt.Sprintf("%d", val)
	default:
		return fmt.Sprintf("%d", val)
	}
}
