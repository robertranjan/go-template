package lib

import (
	"fmt"
	"reflect"
	"strconv"
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

// FixFloatPrecision ajust the number of decimal points for a float64
func FixFloatPrecision(f float64, p int) (n float64, err error) {
	formatString := fmt.Sprintf("%%.%df", p)
	strNum := fmt.Sprintf(formatString, f)
	n, err = strconv.ParseFloat(strNum, 64)
	if err != nil {
		return n, fmt.Errorf("parsing float64 from %q failed: %v", strNum, err)
	}
	return n, nil
}

// IfaceToStrSlice converts []interface to []string
func IfaceToStrSlice(ifaceSlice []interface{}) (str []string) {
	reflectVal := reflect.ValueOf(ifaceSlice)
	for i := 0; i < reflectVal.Len(); i++ {
		str = append(str, fmt.Sprintf("%v", reflectVal.Index(i)))
	}
	return str
}
