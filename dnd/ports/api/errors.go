package api

import (
	"errors"
)

var (
	ErrNotFound          = errors.New("not found")
	ErrRepositoryFailure = errors.New("repository failure")
	ErrInternalError     = errors.New("internal error")
)
