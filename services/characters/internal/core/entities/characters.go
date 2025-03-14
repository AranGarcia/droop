package entities

import (
	"time"
)

// Base attributes shared among entities.
type Base struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Character struct {
	Base
	Class         ClassName    `json:"class" validate:"required,oneof=barbarian bard cleric druid fighter monk paladin ranger rogue sorcerer warlock wizard druid"`
	Level         Level        `json:"level" validate:"required,gte=1,lte=20"`
	Name          string       `json:"name" validate:"required"`
	MaxHealth     int          `json:"max_health" validate:"required,gt=0"`
	CurrentHealth int          `json:"current_health" validate:"ltefield=MaxHealth"`
	TempHealth    int          `json:"temp_health"`
	ArmorClass    int          `json:"armor_class"`
	Strength      AbilityScore `json:"strength" validate:"required"`
	Dexterity     AbilityScore `json:"dexterity" validate:"required"`
	Constitution  AbilityScore `json:"constitution" validate:"required"`
	Intelligence  AbilityScore `json:"intelligence" validate:"required"`
	Wisdom        AbilityScore `json:"wisdom" validate:"required"`
	Charisma      AbilityScore `json:"charisma" validate:"required"`
	Abilities     Abilities    `json:"abilities,omitempty" validate:"dive,keys,oneof=acrobatics animal_handling arcana athletics deception history insight intimidation investigation medicine nature perception performance persuasion religion sleight_of_hand stealth survival,endkeys"`
}

// Validate the character's fields.
func (c Character) Validate() error {
	return validate.Struct(c)
}
