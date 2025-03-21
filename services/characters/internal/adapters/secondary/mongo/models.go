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
	Base          `bson:",inline"`
	Class         string   `bson:"class"`
	Level         int      `bson:"level"`
	Name          string   `bson:"name"`
	MaxHealth     int      `bson:"max_health"`
	CurrentHealth int      `bson:"current_health"`
	TempHealth    int      `bson:"temp_health"`
	ArmorClass    int      `bson:"armor_class"`
	Strength      int      `bson:"strength"`
	Dexterity     int      `bson:"dexterity"`
	Constitution  int      `bson:"constitution"`
	Intelligence  int      `bson:"intelligence"`
	Wisdom        int      `bson:"wisdom"`
	Charisma      int      `bson:"charisma"`
	Proficiencies []string `bson:"proficiencies"`
}

func NewCharacterFromEntity(entity entities.Character) Character {
	character := Character{
		Base: Base{
			ID:        entity.ID,
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
			DeletedAt: entity.DeletedAt,
		},
		Class: string(entity.Class),
		Level: int(entity.Level),
		Name:  entity.Name,

		MaxHealth:     entity.MaxHealth,
		CurrentHealth: entity.CurrentHealth,
		TempHealth:    entity.TempHealth,
		ArmorClass:    entity.ArmorClass,
		Strength:      int(entity.Strength),
		Dexterity:     int(entity.Dexterity),
		Constitution:  int(entity.Constitution),
		Intelligence:  int(entity.Intelligence),
		Wisdom:        int(entity.Wisdom),
		Charisma:      int(entity.Charisma),
		Proficiencies: make([]string, len(entity.Proficiencies)),
	}

	for i, v := range entity.Proficiencies {
		character.Proficiencies[i] = string(v)
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
		Class:         entities.ClassName(c.Class),
		Level:         entities.Level(c.Level),
		Name:          c.Name,
		MaxHealth:     c.MaxHealth,
		CurrentHealth: c.CurrentHealth,
		TempHealth:    c.TempHealth,
		ArmorClass:    c.ArmorClass,
		Strength:      entities.AbilityScore(c.Strength),
		Dexterity:     entities.AbilityScore(c.Dexterity),
		Constitution:  entities.AbilityScore(c.Constitution),
		Intelligence:  entities.AbilityScore(c.Intelligence),
		Wisdom:        entities.AbilityScore(c.Wisdom),
		Charisma:      entities.AbilityScore(c.Charisma),
		Proficiencies: make([]entities.Skill, len(c.Proficiencies)),
	}

	for i, v := range c.Proficiencies {
		entity.Proficiencies[i] = entities.Skill(v)
	}

	return entity
}

func (c *Character) CreatedAtNow() { c.CreatedAt = time.Now() }

func (c *Character) UpdatedAtNow() { c.UpdatedAt = time.Now() }
