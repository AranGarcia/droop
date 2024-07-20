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

type InvalidRequestError struct {
	Fields map[string]string
}

func NewInvalidRequestError(validationError validator.ValidationErrors) InvalidRequestError {
	fields := make(map[string]string, len(validationError))
	for _, fieldError := range validationError {
		fields[fieldError.Field()] = fieldError.Error()
	}
	return InvalidRequestError{Fields: fields}
}

func (r InvalidRequestError) Error() string {
	fields := make([]string, len(r.Fields))
	i := 0
	for field := range r.Fields {
		fields[i] = field
		i++
	}
	return fmt.Sprint("the following fields have invalid values: ", strings.Join(fields, ", "))
}
