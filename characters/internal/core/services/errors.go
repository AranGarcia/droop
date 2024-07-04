package services

import (
	"errors"

	"github.com/AranGarcia/droop/characters/internal/adapters/secondary/mongo"
	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

func repositoryErrorToAPI(err error) error {
	switch err {
	case mongo.ErrorNotFound:
		return api.ErrNotFound
	case mongo.ErrInternalError:
		return errors.Join(api.ErrRepositoryFailure, err)
	}
	return api.ErrInternalError
}
