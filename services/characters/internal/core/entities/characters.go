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
	Level         Level     `json:"level" validate:"required,gte=1,lte=20"`
	Name          string    `json:"name" validate:"required"`
	MaxHealth     int       `json:"max_health" validate:"required,gt=0"`
	CurrentHealth int       `json:"current_health" validate:"ltefield=MaxHealth"`
	TempHealth    int       `json:"temp_health"`
	// ArmorClass protects the character from attacks.
	// Deprecated: Use Armor and the Dexterity modifier instead.
	ArmorClass    int          `json:"armor_class"`
	Strength      AbilityScore `json:"strength" validate:"required"`
	Dexterity     AbilityScore `json:"dexterity" validate:"required"`
	Constitution  AbilityScore `json:"constitution" validate:"required"`
	Intelligence  AbilityScore `json:"intelligence" validate:"required"`
	Wisdom        AbilityScore `json:"wisdom" validate:"required"`
	Charisma      AbilityScore `json:"charisma" validate:"required"`
	Proficiencies []Skill      `json:"proficiencies,omitempty" validate:"dive,oneof=acrobatics animal_handling arcana athletics deception history insight intimidation investigation medicine nature perception performance persuasion religion sleight_of_hand stealth survival"`

	Armor  *Armor `json:"armor"`
	Shield bool   `json:"shield"`
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
		Proficiencies: c.Proficiencies,
	}
}

// Validate the character's fields.
func (c Character) Validate() error {
	return validate.Struct(c)
}

func (c Character) CalculateArmorClass() int {
	var shieldBonus int
	if c.Shield {
		shieldBonus = 2
	}
	hasUnarmoredDefenseFeat, unarmoredDefenseAC := c.unarmoredDefenseFeat(shieldBonus)
	if hasUnarmoredDefenseFeat {
		return unarmoredDefenseAC
	}

	if c.Armor != nil {
		acBonus := c.Armor.CalculateArmorBonus(c)
		return acBonus + shieldBonus
	}

	return 10 + c.Dexterity.Modifier()
}

func (c Character) unarmoredDefenseFeat(shieldBonus int) (bool, int) {
	switch c.Class {
	case BarbarianClass:
		return true, 10 + c.Dexterity.Modifier() + c.Constitution.Modifier() + shieldBonus
	case MonkClass:
		return !c.Shield, 10 + c.Dexterity.Modifier() + c.Wisdom.Modifier()
	default:
		return false, 0
	}
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
