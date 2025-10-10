package commonerrors

import "errors"

var (
	ErrInvalidID         = errors.New("invalid ID")
	ErrNotFound          = errors.New("not found")
	ErrRepositoryFailure = errors.New("repository failure")
	ErrInternalError     = errors.New("internal error")
)
