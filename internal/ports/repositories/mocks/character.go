package mocks

import (
	"context"

	"github.com/AranGarcia/droop/internal/core/entities"
	"github.com/AranGarcia/droop/internal/ports/repositories"
)

// CharacterMockRepository is an in-memory mock implementation for Characters.
type CharacterMockRepository struct {
	characters map[string]entities.Character
}

// Create a Character in memory.
func (c CharacterMockRepository) Create(_ context.Context, e entities.Character) error {
	c.characters[e.ID] = e
	return nil
}

// Retrieve a Character from memory.
func (c CharacterMockRepository) Retrieve(_ context.Context, id string) (*entities.Character, error) {
	character, exists := c.characters[id]
	if !exists {
		return nil, repositories.ErrNotFound
	}
	return &character, nil
}

// Update a Character partially in memory. It returns the new state of the updated resource.
func (c CharacterMockRepository) Update(_ context.Context, id string, in repositories.UpdateCharacterInput) (*entities.Character, error) {
	character, exists := c.characters[id]
	if !exists {
		return nil, repositories.ErrNotFound
	}

	if in.Name != nil {
		character.Name = *in.Name
	}
	if in.Level != nil {
		character.Level = *in.Level
	}
	if in.HitPoints != nil {
		character.HitPoints = *in.HitPoints
	}
	c.characters[id] = character

	return &character, nil
}

// Delete a Character from the memory repository.
func (c CharacterMockRepository) Delete(_ context.Context, id string) error {
	_, exists := c.characters[id]
	if !exists {
		return repositories.ErrNotFound
	}
	delete(c.characters, id)
	return nil
}
