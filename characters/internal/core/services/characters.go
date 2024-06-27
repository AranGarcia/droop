package services

import (
	"context"
	"errors"

	"github.com/AranGarcia/droop/characters/internal/ports/api"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

type Characters struct {
	Repository repositories.Characters
}

func (c Characters) Create(_ context.Context, _ api.CreateCharacterRequest) (*api.CreateCharacterResponse, error) {
	return nil, errors.New("not implemented")
}

func (c Characters) Retrieve(_ context.Context, _ api.RetrieveCharacterRequest) (*api.RetrieveCharacterResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (c Characters) Update(_ context.Context, _ api.UpdateCharacterRequest) (*api.UpdateCharacterResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (c Characters) Delete(_ context.Context, _ api.DeleteCharacterRequest) (*api.DeleteCharacterResponse, error) {
	panic("not implemented") // TODO: Implement
}
