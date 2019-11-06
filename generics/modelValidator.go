package generics

import (
	"fmt"
	"juno/interfaces"
	"reflect"
)

//GenericModelInstanceValidator implements a generic validator for model instances
func GenericModelInstanceValidator(model interfaces.Model) []string {
	var errorData []string = []string{}
	modelType := reflect.TypeOf(model).Elem()
	modelValue := reflect.ValueOf(model).Elem()
	for i := 0; i < modelValue.NumField(); i++ {
		fieldValue := modelValue.Field(i).Interface()
		if fieldValue == "" || fieldValue == 0 {
			errorData = append(errorData, fmt.Sprintf("%s: field is required", modelType.Field(i).Name))
		}
	}
	return errorData
}
