package utils

import (
	"errors"
	"reflect"
)

func StructToMap(obj interface{}) (map[string]any, error) {
	result := make(map[string]any)
	v := reflect.ValueOf(obj)

	// Dereference the pointer if it is a pointer
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure the input is a struct
	if v.Kind() != reflect.Struct {
		return nil, errors.New("object is not a struct")
	}

	t := v.Type() // Get the type after dereferencing

	// Iterate over struct fields
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Check if the field is a pointer and if it is nil
		if value.Kind() == reflect.Ptr && value.IsNil() {
			continue
		}

		// Dereference pointers to get the actual value
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		// Add fields to the map
		result[field.Name] = value.Interface()
	}

	return result, nil
}