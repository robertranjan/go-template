package lib

import (
	"fmt"
	"reflect"
)

func FlattenStruct(obj interface{}, prefix string) map[string]string {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()

	if objType.Kind() != reflect.Struct {
		panic("Invalid input. Expected a struct.")
	}

	result := make(map[string]string)

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := objValue.Field(i)

		fullFieldName := prefix + field.Name

		switch fieldValue.Kind() {
		case reflect.Struct:
			subMap := FlattenStruct(fieldValue.Interface(), fullFieldName+".")
			for key, value := range subMap {
				result[key] = value
			}
		default:
			if fieldValue.CanInterface() {
				result[fullFieldName] = fmt.Sprintf("%v", fieldValue.Interface())
			}
		}
	}

	return result
}
