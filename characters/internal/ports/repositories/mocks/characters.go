package mocks

import (
	"context"
	"errors"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
)

type Characters struct {
	characters map[string]entities.Character
}

func (c *Characters) Create(_ context.Context, character entities.Character) error {
	if _, ok := c.characters[character.ID]; ok {
		return errors.New("duplicate entity error") // TODO: use custom errors
	}
	c.characters[character.ID] = character
	return nil
}

func (c *Characters) Retrieve(_ context.Context, id string) (*entities.Character, error) {
	character, ok := c.characters[id]
	if !ok {
		return nil, errors.New("not found") // TODO: use custom errors
	}
	return &character, nil
}

func (c *Characters) Update(_ context.Context, _ entities.Character) error {
	panic("not implemented") // TODO: Implement
}

func (c *Characters) Delete(_ context.Context, id string) error {
	delete(c.characters, id)
	return nil
}
