package repositories

import (
	"context"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
)

type Characters interface {
	Create(context.Context, entities.Character) error
	Retrieve(context.Context, string) (*entities.Character, error)
	Update(context.Context, entities.Character) error
	Delete(context.Context, string) error
}
