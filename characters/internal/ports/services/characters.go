package services

import (
	"context"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
)

// Characters is an API port for the core domain.
type Characters interface {
	Create(context.Context, CreateCharacterRequest) (CreateCharacterResponse, error)
	Retrieve(context.Context, RetrieveCharacterRequest) (RetrieveCharacterResponse, error)
	Update(context.Context, UpdateCharacterRequest) (UpdateCharacterResponse, error)
	Delete(context.Context, DeleteCharacterRequest) (DeleteCharacterResponse, error)
}

type CreateCharacterRequest struct {
	Character entities.Character
}

type CreateCharacterResponse struct{}

type RetrieveCharacterRequest struct{}

type RetrieveCharacterResponse struct{}

type UpdateCharacterRequest struct{}

type UpdateCharacterResponse struct{}

type DeleteCharacterRequest struct{}

type DeleteCharacterResponse struct{}
