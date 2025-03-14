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
	Proficiencies []Skill      `json:"proficiencies,omitempty" validate:"dive,oneof=acrobatics animal_handling arcana athletics deception history insight intimidation investigation medicine nature perception performance persuasion religion sleight_of_hand stealth survival"`
}

// Validate the character's fields.
func (c Character) Validate() error {
	return validate.Struct(c)
}

// // CalculateArrmorClass uses the character's inner state from the armor, class feats, and shield
// // to calculate the total AC.
// func (c Character) CalculateArmorClass() int {
// 	var shieldBonus int
// 	if c.Shield {
// 		shieldBonus = 2
// 	}
// 	hasUnarmoredDefenseFeat, unarmoredDefenseAC := c.unarmoredDefenseFeat(shieldBonus)
// 	if hasUnarmoredDefenseFeat {
// 		return unarmoredDefenseAC
// 	}

// 	if c.Armor != nil {
// 		acBonus := c.Armor.CalculateArmorBonus(c)
// 		return acBonus + shieldBonus
// 	}

// 	return 10 + c.Dexterity.Modifier()
// }

// // unarmoredDefenseFeat calculates the character's unarmored defense status according to the class.
// // If it applies, then it returns true plus the sum of appropriate modifiers to calculate the AC
// // bonus.
// func (c Character) unarmoredDefenseFeat(shieldBonus int) (bool, int) {
// 	switch c.Class {
// 	case BarbarianClass:
// 		return true, 10 + c.Dexterity.Modifier() + c.Constitution.Modifier() + shieldBonus
// 	case MonkClass:
// 		return !c.Shield, 10 + c.Dexterity.Modifier() + c.Wisdom.Modifier()
// 	default:
// 		return false, 0
// 	}
// }
