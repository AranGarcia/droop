package entities

import "time"

// Base attributes shared among entities.
type Base struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Character struct {
	Base
	Class         ClassName `json:"class" validate:"required,oneof=barbarian bard cleric druid fighter monk paladin ranger rogue sorcerer warlock wizard druid"`
	Level         int       `json:"level" validate:"required,gte=1,lte=20"`
	Name          string    `json:"name" validate:"required"`
	MaxHealth     int       `json:"max_health" validate:"required,gt=0"`
	CurrentHealth int       `json:"current_health" validate:"ltefield=MaxHealth"`
	TempHealth    int       `json:"temp_health"`
	ArmorClass    int       `json:"armor_class"`
	Strength      int       `json:"strength" validate:"required"`
	Dexterity     int       `json:"dexterity" validate:"required"`
	Constitution  int       `json:"constitution" validate:"required"`
	Intelligence  int       `json:"intelligence" validate:"required"`
	Wisdom        int       `json:"wisdom" validate:"required"`
	Charisma      int       `json:"charisma" validate:"required"`
	Proficiencies []Skill   `json:"proficiencies" validate:"oneof=dive,acrobatics animal_handling arcana athletics deception history insight intimidation investigation medicine nature perception performance persuasion religion sleight_of_hand stealth survival"`
}

// Copy creates a deep copy of the character.
func (c Character) Copy() Character {
	return Character{
		Base: Base{
			ID:        c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			DeletedAt: copyTimePtr(c.DeletedAt),
		},
		Class:         c.Class,
		Level:         c.Level,
		Name:          c.Name,
		MaxHealth:     c.MaxHealth,
		CurrentHealth: c.CurrentHealth,
		TempHealth:    c.TempHealth,
		ArmorClass:    c.ArmorClass,
		Strength:      c.Strength,
		Dexterity:     c.Dexterity,
		Constitution:  c.Constitution,
		Intelligence:  c.Intelligence,
		Wisdom:        c.Wisdom,
		Charisma:      c.Charisma,
	}
}

// Validate the character's fields.
func (c Character) Validate() error {
	return validate.Struct(c)
}

func copyTimePtr(t *time.Time) *time.Time {
	if t == nil {
		return nil
	}
	t2 := time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
		t.Location(),
	)
	return &t2
}
