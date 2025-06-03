package middleware

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s any) []string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errs := err.(validator.ValidationErrors)
	details := make([]string, 0, len(errs))

	valType := reflect.TypeOf(s)
	if valType.Kind() == reflect.Ptr {
		valType = valType.Elem()
	}

	for _, e := range errs {
		structFieldName := e.StructField()

		if field, found := valType.FieldByName(structFieldName); found {
			jsonTag := field.Tag.Get("json")
			jsonField := strings.Split(jsonTag, ",")[0]
			if jsonField == "" {
				jsonField = strings.ToLower(structFieldName)
			}
			details = append(details, fmt.Sprintf("'%s' need %s", jsonField, e.Tag()))
		} else {
			lowerName := strings.ToLower(structFieldName)
			details = append(details, fmt.Sprintf("'%s' need %s", lowerName, e.Tag()))
		}
	}
	return details
}
