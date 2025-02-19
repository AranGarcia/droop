package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

type UpdateCharacterRequest struct {
	ID           string              `json:"id"`
	Class        *entities.ClassName `json:"class,omitempty"`
	Level        *int                `json:"level,omitempty"`
	Name         *string             `json:"name,omitempty"`
	HealthPoints *int                `json:"health_points,omitempty"`
	ArmorClass   *int                `json:"armor_class,omitempty"`
	Strength     *int                `json:"strength,omitempty"`
	Dexterity    *int                `json:"dexterity,omitempty"`
	Constitution *int                `json:"constitution,omitempty"`
	Intelligence *int                `json:"intelligence,omitempty"`
	Wisdom       *int                `json:"wisdom,omitempty"`
	Charisma     *int                `json:"charisma,omitempty"`
}

type UpdateCharacterResponse struct {
	Character entities.Character
}
