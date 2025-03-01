package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

// CreateCharacterRequest is the input for Characters.Create.
type CreateCharacterRequest struct {
	Level        int                `json:"level"`
	Class        entities.ClassName `json:"class"`
	Name         string             `json:"name"`
	MaxHealth    int                `json:"max_health"`
	ArmorClass   int                `json:"armor_class"`
	Strength     int                `json:"strength"`
	Dexterity    int                `json:"dexterity"`
	Constitution int                `json:"constitution"`
	Intelligence int                `json:"intelligence"`
	Wisdom       int                `json:"wisdom"`
	Charisma     int                `json:"charisma"`
}

// CreateCharacterResponse is the output for Characters.Create.
type CreateCharacterResponse struct {
	Character entities.Character
}
