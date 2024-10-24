package mock

import (
	"context"
	"errors"
	"time"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

type Characters struct {
	Entities map[string]entities.Character
}

func NewMockCharacterService() *Characters {
	return &Characters{
		Entities: make(map[string]entities.Character),
	}
}

func (c *Characters) Create(_ context.Context, request api.CreateCharacterRequest) (*api.CreateCharacterResponse, error) {
	id := time.Now().String()
	character := entities.Character{
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
	c.Entities[id] = character
	response := api.CreateCharacterResponse{Character: character}
	return &response, nil
}

func (c *Characters) Retrieve(_ context.Context, request api.RetrieveCharacterRequest) (*api.RetrieveCharacterResponse, error) {
	character, ok := c.Entities[request.ID]
	if !ok {
		return nil, api.ErrNotFound
	}
	response := &api.RetrieveCharacterResponse{Character: character}
	return response, nil
}

func (c *Characters) Update(_ context.Context, request api.UpdateCharacterRequest) (*api.UpdateCharacterResponse, error) {
	character, ok := c.Entities[request.ID]
	if !ok {
		return nil, api.ErrNotFound
	}
	if request.Level != nil {
		character.Level = *request.Level
	}
	if request.Name != nil {
		character.Name = *request.Name
	}
	if request.HealthPoints != nil {
		character.HealthPoints = *request.HealthPoints
	}
	if request.ArmorClass != nil {
		character.ArmorClass = *request.ArmorClass
	}
	if request.Strength != nil {
		character.Strength = *request.Strength
	}
	if request.Dexterity != nil {
		character.Dexterity = *request.Dexterity
	}
	if request.Constitution != nil {
		character.Constitution = *request.Constitution
	}
	if request.Intelligence != nil {
		character.Intelligence = *request.Intelligence
	}
	if request.Wisdom != nil {
		character.Wisdom = *request.Wisdom
	}
	if request.Charisma != nil {
		character.Charisma = *request.Charisma
	}
	c.Entities[request.ID] = character
	response := &api.UpdateCharacterResponse{Character: character}
	return response, nil
}

func (c *Characters) Delete(_ context.Context, request api.DeleteCharacterRequest) (*api.DeleteCharacterResponse, error) {
	if _, ok := c.Entities[request.ID]; !ok {
		return nil, api.ErrNotFound
	}
	delete(c.Entities, request.ID)
	response := &api.DeleteCharacterResponse{}
	return response, nil
}

func (c *Characters) List(_ context.Context, request api.ListCharactersRequest) (*api.ListCharactersResponse, error) {
	return nil, errors.New("not implemented")
}
