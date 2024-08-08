package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

type ListCharactersRequest struct{}

type ListCharactersResponse struct {
	Characters []entities.Character `json:"characters"`
}
