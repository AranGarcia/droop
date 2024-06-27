package services

import "errors"

var (
	ErrRepositoryFailure = errors.New("repository failure")
	ErrInternalError     = errors.New("internal error")
)

func repositoryErrorToAPI(_ error) error {
	return ErrInternalError
}
