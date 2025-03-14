package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

// CreateCharacterRequest is the input for Characters.Create.
type CreateCharacterRequest struct {
	Level        entities.Level        `json:"level"`
	Class        entities.ClassName    `json:"class"`
	Name         string                `json:"name"`
	MaxHealth    int                   `json:"max_health"`
	ArmorClass   int                   `json:"armor_class"`
	Strength     entities.AbilityScore `json:"strength"`
	Dexterity    entities.AbilityScore `json:"dexterity"`
	Constitution entities.AbilityScore `json:"constitution"`
	Intelligence entities.AbilityScore `json:"intelligence"`
	Wisdom       entities.AbilityScore `json:"wisdom"`
	Charisma     entities.AbilityScore `json:"charisma"`
	Abilities    entities.Abilities    `json:"abilities"`
}

// CreateCharacterResponse is the output for Characters.Create.
type CreateCharacterResponse struct {
	Character entities.Character
}
