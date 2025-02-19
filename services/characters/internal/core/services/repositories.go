package services

import (
	"github.com/AranGarcia/droop/characters/internal/core/api"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

func getRepositoryUpdateFields(request api.UpdateCharacterRequest) repositories.CharacterFields {
	fields := repositories.CharacterFields{
		Level:        request.Level,
		Name:         request.Name,
		HealthPoints: request.HealthPoints,
		ArmorClass:   request.ArmorClass,
		Strength:     request.Strength,
		Dexterity:    request.Dexterity,
		Constitution: request.Constitution,
		Intelligence: request.Intelligence,
		Wisdom:       request.Wisdom,
		Charisma:     request.Charisma,
	}
	return fields
}
