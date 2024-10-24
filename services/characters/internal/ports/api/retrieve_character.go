package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

type RetrieveCharacterRequest struct {
	ID string `json:"id"`
}

type RetrieveCharacterResponse struct {
	Character entities.Character
}
