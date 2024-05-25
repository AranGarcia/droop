package repositories

import (
	"context"

	"github.com/AranGarcia/droop/internal/core/entities"
)

type Character interface {
	Create(context.Context, entities.Character) error
	Retrieve(context.Context, string) (*entities.Character, error)
	Update(context.Context, string, UpdateCharacterInput) (*entities.Character, error)
	Delete(context.Context, string) error
}

// UpdateCharacterInput is the input to Character.Update. Any non-nil field will be updated.
type UpdateCharacterInput struct {
	Name      *string
	Level     *uint
	HitPoints *int
}
