package services

import (
	"context"

	"github.com/AranGarcia/droop/internal/core/entities"
)

type Character struct {
}

func (c Character) Create(_ context.Context, _ entities.Character) error {
	panic("not implemented") // TODO: Implement
}

// Get a Character by its ID.
func (c Character) Get(_ context.Context, _ string) (entities.Character, error) {
	panic("not implemented") // TODO: Implement
}

// Update a Character partially
func (c Character) Update(_ context.Context) {
	panic("not implemented") // TODO: Implement
}

// Delete a Character by its ID.
func (c Character) Delete(_ context.Context, _ string) {
	panic("not implemented") // TODO: Implement
}
