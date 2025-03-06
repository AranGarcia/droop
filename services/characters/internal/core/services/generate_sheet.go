package services

import (
	"context"

	"github.com/AranGarcia/droop/characters/internal/core/api"
)

func (c Characters) GenerateSheet(ctx context.Context, request api.GenerateSheetRequest) (*api.GenerateSheetResponse, error) {
	character, err := c.repository.Retrieve(ctx, request.ID)
	if err != nil {
		return nil, repositoryErrorToAPI(err)
	}
	response := &api.GenerateSheetResponse{
		Characer:    *character,
		Proficiency: character.Level.CalculateProficiencyBonus(),
		Initiative:  character.Dexterity.Modifier(),
	}
	return response, nil
}
