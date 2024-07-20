package services

import (
	"errors"

	"github.com/AranGarcia/droop/characters/internal/adapters/secondary/mongo"
	"github.com/AranGarcia/droop/characters/internal/ports/api"
	"github.com/go-playground/validator/v10"
)

func repositoryErrorToAPI(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return api.NewInvalidRequestError(validationErrors)
	}

	switch err {
	case mongo.ErrorNotFound:
		return api.ErrNotFound
	case mongo.ErrInternalError:
		return errors.Join(api.ErrRepositoryFailure, err)
	}
	return api.ErrInternalError
}
