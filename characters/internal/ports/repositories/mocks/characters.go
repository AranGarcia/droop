package mocks

import (
	"context"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

type Characters struct {
	inMemory map[string]entities.Character
}

func (c *Characters) Create(_ context.Context, character entities.Character) error {
	if _, ok := c.inMemory[character.ID]; ok {
		return repositories.ErrDuplicateEntity
	}
	c.inMemory[character.ID] = character
	return nil
}

func (c *Characters) Retrieve(_ context.Context, id string) (*entities.Character, error) {
	character, ok := c.inMemory[id]
	if !ok {
		return nil, repositories.ErrNotFound
	}
	return &character, nil
}

func (c *Characters) Update(ctx context.Context, id string, fields repositories.CharacterFields) (*entities.Character, error) {
	character, ok := c.inMemory[id]
	if !ok {
		return nil, repositories.ErrNotFound
	}
	if fields.Level != nil {
		character.Level = *fields.Level
	}
	if fields.Name != nil {
		character.Name = *fields.Name
	}
	if fields.HealthPoints != nil {
		character.HealthPoints = *fields.HealthPoints
	}
	if fields.ArmorClass != nil {
		character.ArmorClass = *fields.ArmorClass
	}
	if fields.Strength != nil {
		character.Strength = *fields.Strength
	}
	if fields.Dexterity != nil {
		character.Dexterity = *fields.Dexterity
	}
	if fields.Constitution != nil {
		character.Constitution = *fields.Constitution
	}
	if fields.Intelligence != nil {
		character.Intelligence = *fields.Intelligence
	}
	if fields.Wisdom != nil {
		character.Wisdom = *fields.Wisdom
	}
	if fields.Charisma != nil {
		character.Charisma = *fields.Charisma
	}

	return &character, nil
}

func (c *Characters) Delete(_ context.Context, id string) error {
	delete(c.inMemory, id)
	return nil
}
