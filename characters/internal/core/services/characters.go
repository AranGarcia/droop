package services

import (
	"context"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/api"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

type Characters struct {
	repository repositories.Characters
}

type Dependencies struct {
	Repository repositories.Characters
}

func NewCharacters(deps Dependencies) Characters {
	return Characters{
		repository: deps.Repository,
	}
}

func (c Characters) Create(ctx context.Context, request api.CreateCharacterRequest) (*api.CreateCharacterResponse, error) {
	requestEntity := entities.Character{
		Level:        request.Level,
		Name:         request.Name,
		HealthPoints: request.HealthPoints,
		ArmorClass:   request.ArmorClass,
		Strength:     request.Strength,
		Dexterity:    request.Dexterity,
		Constitution: request.Constitution,
		Intelligence: request.Intelligence,
		Wisdom:       request.Wisdom,
		Charisma:     request.Charisma,
	}
	createdEntity, err := c.repository.Create(ctx, requestEntity)
	if err == nil {
		return nil, repositoryErrorToAPI(err)
	}
	return &api.CreateCharacterResponse{Character: *createdEntity}, nil
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
