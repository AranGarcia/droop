package repositories

import "errors"

var (
	ErrDuplicateEntity = errors.New("duplicate entity")
	ErrNotFound        = errors.New("not found")
)
