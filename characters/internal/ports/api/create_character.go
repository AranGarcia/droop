package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

// CreateCharacterRequest is the input for Characters.Create.
type CreateCharacterRequest struct {
	Level        int    `json:"level"`
	Name         string `json:"name"`
	HealthPoints int    `json:"health_points"`
	ArmorClass   int    `json:"armor_class"`
	Strength     int    `json:"strength"`
	Dexterity    int    `json:"dexterity"`
	Constitution int    `json:"constitution"`
	Intelligence int    `json:"intelligence"`
	Wisdom       int    `json:"wisdom"`
	Charisma     int    `json:"charisma"`
}

// CreateCharacterResponse is the output for Characters.Create.
type CreateCharacterResponse struct {
	Character entities.Character
}
