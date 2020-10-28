package recursion

import "reflect"

// Recursion is
func recursion(input interface{}, fn func(entry string)) {
	value := extractValue(input)

	runnerValue := func(value reflect.Value) {
		recursion(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			runnerValue(value.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			runnerValue(value.Index(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			runnerValue(value.MapIndex(key))
		}
	}
}

func extractValue(input interface{}) reflect.Value {
	value := reflect.ValueOf(input)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
