package api

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrNotFound          = errors.New("not found")
	ErrRepositoryFailure = errors.New("repository failure")
	ErrInternalError     = errors.New("internal error")
)

// InvalidRequestError contains information on the invalid request that was made to the service.
type InvalidRequestError struct {
	// Fields of the request that contain invalid values. Each field name is mapped to the reason.
	Fields map[string]string
}

// NewInvalidRequestError is created from validating at the struct level.
func NewInvalidRequestError(validationError validator.ValidationErrors) InvalidRequestError {
	fields := make(map[string]string, len(validationError))
	for _, fieldError := range validationError {
		fields[fieldError.Field()] = fieldError.Error()
	}
	return InvalidRequestError{Fields: fields}
}

// Error implements the error interface.
func (r InvalidRequestError) Error() string {
	fields := make([]string, len(r.Fields))
	i := 0
	for field := range r.Fields {
		fields[i] = field
		i++
	}
	return fmt.Sprint("the following fields have invalid values: ", strings.Join(fields, ", "))
}
