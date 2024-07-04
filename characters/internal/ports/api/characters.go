package api

import (
	"context"
)

type Characters interface {
	Create(context.Context, CreateCharacterRequest) (*CreateCharacterResponse, error)
	Retrieve(context.Context, RetrieveCharacterRequest) (*RetrieveCharacterResponse, error)
	Update(context.Context, UpdateCharacterRequest) (*UpdateCharacterResponse, error)
	Delete(context.Context, DeleteCharacterRequest) (*DeleteCharacterResponse, error)
}

type UpdateCharacterRequest struct{}

type UpdateCharacterResponse struct{}

type DeleteCharacterRequest struct{}

type DeleteCharacterResponse struct{}
