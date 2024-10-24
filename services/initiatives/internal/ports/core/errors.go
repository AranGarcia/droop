package core

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

// ExternalError are failures caused by external causes.
type ExternalError struct {
	Message string
	Err     error
}

// Error implements the error interface.
func (r ExternalError) Error() string {
	return r.Message
}
