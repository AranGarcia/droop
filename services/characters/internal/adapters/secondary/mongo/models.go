package mongo

import (
	"time"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
)

// Base attributes shared among entities.
type Base struct {
	ID        string     `bson:"_id,omitempty"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

type Character struct {
	Base         `bson:",inline"`
	Class        string `bson:"class"`
	Level        int    `bson:"level"`
	Name         string `bson:"name"`
	HealthPoints int    `bson:"health_points"`
	ArmorClass   int    `bson:"armor_class"`
	Strength     int    `bson:"strength"`
	Dexterity    int    `bson:"dexterity"`
	Constitution int    `bson:"constitution"`
	Intelligence int    `bson:"intelligence"`
	Wisdom       int    `bson:"wisdom"`
	Charisma     int    `bson:"charisma"`
}

func NewCharacterFromEntity(entity entities.Character) Character {
	character := Character{
		Base: Base{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
			DeletedAt: entity.DeletedAt,
		},
		Class:        string(entity.Class),
		Level:        entity.Level,
		Name:         entity.Name,
		HealthPoints: entity.HealthPoints,
		ArmorClass:   entity.ArmorClass,
		Strength:     entity.Strength,
		Dexterity:    entity.Dexterity,
		Constitution: entity.Constitution,
		Intelligence: entity.Intelligence,
		Wisdom:       entity.Wisdom,
		Charisma:     entity.Charisma,
	}
	return character
}

func (c *Character) ToEntity() entities.Character {
	entity := entities.Character{
		Base: entities.Base{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			DeletedAt: c.DeletedAt,
		},
		Class:        entities.ClassName(c.Class),
		Level:        c.Level,
		Name:         c.Name,
		HealthPoints: c.HealthPoints,
		ArmorClass:   c.ArmorClass,
		Strength:     c.Strength,
		Dexterity:    c.Dexterity,
		Constitution: c.Constitution,
		Intelligence: c.Intelligence,
		Wisdom:       c.Wisdom,
		Charisma:     c.Charisma,
	}
	return entity
}

func (c *Character) CreatedAtNow() { c.CreatedAt = time.Now() }

func (c *Character) UpdatedAtNow() { c.UpdatedAt = time.Now() }
