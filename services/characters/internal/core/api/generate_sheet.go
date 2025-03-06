package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

type GenerateSheetRequest struct {
	// ID of the character to generate the sheet for.
	ID string
}

type GenerateSheetResponse struct {
	Characer    entities.Character
	Proficiency int
	Initiative  int
}
