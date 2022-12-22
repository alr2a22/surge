package validation

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func Validate(v interface{}) (map[string]string, bool) {
	validationResult := make(map[string]string)
	isValid := true

	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	if errs, ok := validate.Struct(v).(validator.ValidationErrors); ok {
		if len(errs) != 0 {
			isValid = false
			for _, err := range errs {
				validationResult[err.Field()] = err.Tag()
			}
		}
	}
	return validationResult, isValid
}
