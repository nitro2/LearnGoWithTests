package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fmt.Printf("test: %v\n", field)
		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
