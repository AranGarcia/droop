package repositories

import (
	"context"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
)

type Characters interface {
	// Create a character.
	Create(context.Context, entities.Character) (*entities.Character, error)
	// Retrieve a character by its ID.
	Retrieve(context.Context, string) (*entities.Character, error)
	// Update the character that matches the ID with the included fields.
	Update(context.Context, string, CharacterFields) (*entities.Character, error)
	// Delete a character.
	Delete(context.Context, string) error
	// List of characters to be retrieved. The result is paginated with an offset and a limit.
	List(context.Context, int, int) ([]entities.Character, error)
}

// CharacterFields helps define the subset of fields of a entities.Character. Any nil value will
// be ignored.
type CharacterFields struct {
	Level         *int
	Class         *string
	Name          *string
	MaxHealth     *int
	CurrentHealth *int
	TempHealth    *int
	ArmorClass    *int
	Strength      *int
	Dexterity     *int
	Constitution  *int
	Intelligence  *int
	Wisdom        *int
	Charisma      *int
	HealthPoints  *int // TODO: remove
}
