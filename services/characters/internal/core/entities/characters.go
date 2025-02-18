package entities

import "time"

// Base attributes shared among entities.
type Base struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ClassName string

const (
	BarbarianClass ClassName = "barbarian"
	BardClass      ClassName = "bard"
	ClericClass    ClassName = "cleric"
	DruidClass     ClassName = "druid"
	FighterClass   ClassName = "fighter"
	MonkClass      ClassName = "monk"
	PaladinClass   ClassName = "paladin"
	RangerClass    ClassName = "ranger"
	RogueClass     ClassName = "rogue"
	SorcererClass  ClassName = "sorcerer"
	WarlockClass   ClassName = "warlock"
	WizardClass    ClassName = "wizard"
)

type Character struct {
	Base
	Class        ClassName `json:"class" validate:"required,oneof='barbarian bard cleric druid fighter monk paladin ranger rogue sorcerer warlock wizard druid'"`
	Level        int       `json:"level" validate:"required,gte=1,lte=20"`
	Name         string    `json:"name" validate:"required"`
	HealthPoints int       `json:"health_points"`
	ArmorClass   int       `json:"armor_class"`
	Strength     int       `json:"strength" validate:"required"`
	Dexterity    int       `json:"dexterity" validate:"required"`
	Constitution int       `json:"constitution" validate:"required"`
	Intelligence int       `json:"intelligence" validate:"required"`
	Wisdom       int       `json:"wisdom" validate:"required"`
	Charisma     int       `json:"charisma" validate:"required"`
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
		Class:        c.Class,
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
