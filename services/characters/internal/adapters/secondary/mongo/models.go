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

type Skill struct {
	Proficient bool `bson:"proficient"`
	Bonus      int  `bson:"proficient"`
}

func NewSkillFromEntity(entity entities.Skill) Skill {
	return Skill{
		Proficient: entity.Proficient,
		Bonus:      entity.Bonus,
	}
}

func (s Skill) ToEntity() entities.Skill {
	return entities.Skill{
		Proficient: s.Proficient,
		Bonus:      s.Bonus,
	}
}

type Character struct {
	Base          `bson:",inline"`
	Class         string           `bson:"class"`
	Level         int              `bson:"level"`
	Name          string           `bson:"name"`
	MaxHealth     int              `bson:"max_health"`
	CurrentHealth int              `bson:"current_health"`
	TempHealth    int              `bson:"temp_health"`
	ArmorClass    int              `bson:"armor_class"`
	Strength      int              `bson:"strength"`
	Dexterity     int              `bson:"dexterity"`
	Constitution  int              `bson:"constitution"`
	Intelligence  int              `bson:"intelligence"`
	Wisdom        int              `bson:"wisdom"`
	Charisma      int              `bson:"charisma"`
	Abilities     map[string]Skill `bson:"proficiencies"`
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
		Abilities:     make(map[string]Skill, len(entity.Abilities)),
	}

	for k, v := range entity.Abilities {
		character.Abilities[k] = NewSkillFromEntity(v)
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
		Abilities:     make(entities.Abilities, len(c.Abilities)),
	}

	for k, v := range c.Abilities {
		entity.Abilities[k] = v.ToEntity()
	}

	return entity
}

func (c *Character) CreatedAtNow() { c.CreatedAt = time.Now() }

func (c *Character) UpdatedAtNow() { c.UpdatedAt = time.Now() }
