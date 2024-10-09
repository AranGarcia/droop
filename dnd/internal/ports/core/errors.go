package core

import (
	"errors"
	"fmt"
)

var (
	ErrNoIDProvided      = errors.New("no ID provided")
	ErrInvalidInput      = errors.New("invalid input")
	ErrNotFound          = errors.New("not found")
	ErrRepositoryFailure = errors.New("repository failure")
	ErrInternalError     = errors.New("internal error")
)

// InternalError wraps any error that failed internally.
type InternalError struct {
	Message string
	Err     error
}

// Error implements the error interface.
func (i InternalError) Error() string {
	return fmt.Sprintf("%s; %v", i.Message, i.Err)
}

// ExternalServiceError wraps any error that occurred when using an external service.
type ExternalServiceError struct {
	Service string
	Err     error
}

func NewExternalServiceError(service string, err error) ExternalServiceError {
	return ExternalServiceError{
		Service: service,
		Err:     err,
	}
}

// Error inmplements the error interface.
func (ext ExternalServiceError) Error() string {
	return fmt.Sprintf("external error when using service %q; %v", ext.Service, ext.Err)
}
