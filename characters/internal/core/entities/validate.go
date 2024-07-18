package entities

import "github.com/go-playground/validator/v10"

// validate is the singleton that caches struct validation info.
var validate *validator.Validate

func init() {
	validate = validator.New()
}
