package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

type UpdateCharacterRequest struct {
	ID            string              `json:"id"`
	Class         *entities.ClassName `json:"class,omitempty"`
	Level         *entities.Level     `json:"level,omitempty"`
	Name          *string             `json:"name,omitempty"`
	MaxHealth     *int                `json:"max_health,omitempty"`
	CurrentHealth *int                `json:"current_health,omitempty"`
	TempHealth    *int                `json:"temp_health,omitempty"`
	ArmorClass    *int                `json:"armor_class,omitempty"`
	Strength      *int                `json:"strength,omitempty"`
	Dexterity     *int                `json:"dexterity,omitempty"`
	Constitution  *int                `json:"constitution,omitempty"`
	Intelligence  *int                `json:"intelligence,omitempty"`
	Wisdom        *int                `json:"wisdom,omitempty"`
	Charisma      *int                `json:"charisma,omitempty"`
	Proficiencies []entities.Skill    `json:"proficiencies,omitempty"`
}

type UpdateCharacterResponse struct {
	Character entities.Character
}
