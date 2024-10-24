package mocks

import (
	"context"
	"errors"
	"time"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

type Characters struct {
	inMemory map[string]entities.Character
}

// Create a Character in memory. It returns a copy of the created entity with a newly set ID.
func (c *Characters) Create(_ context.Context, character entities.Character) (*entities.Character, error) {
	createdCharacter := character.Copy()
	id := time.Now().Format(time.RFC3339Nano)
	createdCharacter.ID = id
	if _, ok := c.inMemory[id]; ok {
		return nil, repositories.ErrDuplicateEntity
	}
	c.inMemory[createdCharacter.ID] = createdCharacter
	return &createdCharacter, nil
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

func (c *Characters) List(_ context.Context, offset, limit int) ([]entities.Character, error) {
	return nil, errors.New("unimplemented")
}
