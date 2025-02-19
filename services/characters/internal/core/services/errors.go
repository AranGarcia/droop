package services

import (
	"errors"

	"github.com/go-playground/validator/v10"

	"github.com/AranGarcia/droop/characters/internal/adapters/secondary/mongo"
	"github.com/AranGarcia/droop/characters/internal/core/api"
)

func repositoryErrorToAPI(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		return api.NewInvalidRequestError(validationErrors)
	}

	switch err {
	case mongo.ErrInvalidID:
		return api.ErrInvalidID
	case mongo.ErrorNotFound:
		return api.ErrNotFound
	case mongo.ErrInternalError:
		return errors.Join(api.ErrRepositoryFailure, err)
	}
	return api.ErrInternalError
}
