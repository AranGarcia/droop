package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

type UpdateCharacterRequest struct {
	ID            string                 `json:"id"`
	Class         *entities.ClassName    `json:"class,omitempty"`
	Level         *entities.Level        `json:"level,omitempty"`
	Name          *string                `json:"name,omitempty"`
	MaxHealth     *int                   `json:"max_health,omitempty"`
	CurrentHealth *int                   `json:"current_health,omitempty"`
	TempHealth    *int                   `json:"temp_health,omitempty"`
	ArmorClass    *int                   `json:"armor_class,omitempty"`
	Strength      *entities.AbilityScore `json:"strength,omitempty"`
	Dexterity     *entities.AbilityScore `json:"dexterity,omitempty"`
	Constitution  *entities.AbilityScore `json:"constitution,omitempty"`
	Intelligence  *entities.AbilityScore `json:"intelligence,omitempty"`
	Wisdom        *entities.AbilityScore `json:"wisdom,omitempty"`
	Charisma      *entities.AbilityScore `json:"charisma,omitempty"`
	Proficiencies []entities.Skill       `json:"proficiencies,omitempty"`
}

type UpdateCharacterResponse struct {
	Character entities.Character
}
