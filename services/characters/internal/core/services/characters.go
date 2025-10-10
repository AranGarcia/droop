package services

import (
	"context"

	"github.com/AranGarcia/droop/characters/internal/core/api"
	"github.com/AranGarcia/droop/characters/internal/core/entities"
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
		Class:         request.Class,
		Level:         request.Level,
		Name:          request.Name,
		MaxHealth:     request.MaxHealth,
		CurrentHealth: request.MaxHealth, // Set initial current health as the max
		ArmorClass:    request.ArmorClass,
		Strength:      request.Strength,
		Dexterity:     request.Dexterity,
		Constitution:  request.Constitution,
		Intelligence:  request.Intelligence,
		Wisdom:        request.Wisdom,
		Charisma:      request.Charisma,
		Abilities:     request.Abilities,
	}

	if err := requestEntity.Validate(); err != nil {
		return nil, repositoryErrorToAPI(err)
	}

	createdEntity, err := c.repository.Create(ctx, requestEntity)
	if err != nil {
		return nil, repositoryErrorToAPI(err)
	}
	return &api.CreateCharacterResponse{Character: *createdEntity}, nil
}

func (c Characters) Retrieve(ctx context.Context, request api.RetrieveCharacterRequest) (*api.RetrieveCharacterResponse, error) {
	character, err := c.repository.Retrieve(ctx, request.ID)
	if err != nil {
		return nil, repositoryErrorToAPI(err)
	}
	return &api.RetrieveCharacterResponse{Character: *character}, nil
}

func (c Characters) Update(ctx context.Context, request api.UpdateCharacterRequest) (*api.UpdateCharacterResponse, error) {
	repositoryUpdateFields := getRepositoryUpdateFields(request)
	character, err := c.repository.Update(ctx, request.ID, repositoryUpdateFields)
	if err != nil {
		return nil, repositoryErrorToAPI(err)
	}
	return &api.UpdateCharacterResponse{Character: *character}, nil
}

func (c Characters) Delete(ctx context.Context, request api.DeleteCharacterRequest) (*api.DeleteCharacterResponse, error) {
	if err := c.repository.Delete(ctx, request.ID); err != nil {
		return nil, repositoryErrorToAPI(err)
	}
	return &api.DeleteCharacterResponse{}, nil
}

func (c Characters) List(ctx context.Context, request api.ListCharactersRequest) (*api.ListCharactersResponse, error) {
	// Default sort key
	if request.Sort == "" {
		request.Sort = api.CreatedAt
	}

	if err := request.Validate(); err != nil {
		return nil, err
	}
	characters, err := c.repository.List(ctx, request.Offset, request.Size)
	if err != nil {
		return nil, repositoryErrorToAPI(err)
	}
	return &api.ListCharactersResponse{Characters: characters}, nil
}
