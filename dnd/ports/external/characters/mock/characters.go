package mock

import (
	"context"

	"github.con/AranGarcia/droop/dnd/core/entities"
	"github.con/AranGarcia/droop/dnd/ports/api"
)

// Characters is a mock implementation of the external Character service.
type Characters struct {
	// Data is the in-memory storage for the characters.
	Data map[string]entities.Character
}

// Retrieve a character from memory.s
func (c Characters) Retrieve(_ context.Context, id string) (*entities.Character, error) {
	character, ok := c.Data[id]
	if !ok {
		return nil, api.ErrNotFound
	}
	return &character, nil
}
