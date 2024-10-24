package entities

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// validate is the singleton that caches struct validation info.
var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
	})
}
