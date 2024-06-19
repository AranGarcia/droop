package repositories

import (
	"context"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
)

type Characters interface {
	Create(context.Context, entities.Character) error
	Retrieve(context.Context, string) (*entities.Character, error)
	Update(context.Context, string, CharacterFields) (*entities.Character, error)
	Delete(context.Context, string) error
}

// CharacterFields helps define the subset of fields of a entities.Character. Any nil value will
// be ignored.
type CharacterFields struct {
	Level        *int
	Name         *string
	HealthPoints *int
	ArmorClass   *int
	Strength     *int
	Dexterity    *int
	Constitution *int
	Intelligence *int
	Wisdom       *int
	Charisma     *int
}
