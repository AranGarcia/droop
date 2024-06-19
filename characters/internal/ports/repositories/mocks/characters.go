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

func (c *Characters) Update(_ context.Context, _ entities.Character) error {
	panic("not implemented") // TODO: Implement
}

func (c *Characters) Delete(_ context.Context, id string) error {
	delete(c.inMemory, id)
	return nil
}
