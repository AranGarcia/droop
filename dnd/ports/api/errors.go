package api

import (
	"errors"
)

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrNotFound          = errors.New("not found")
	ErrRepositoryFailure = errors.New("repository failure")
	ErrInternalError     = errors.New("internal error")
)
