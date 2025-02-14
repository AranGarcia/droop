package api

import (
	"github.com/AranGarcia/droop/characters/internal/core/entities"
)

const (
	CreatedAt string = "created_at"
	UpdatedAt string = "updated_at"
)

type ListCharactersRequest struct {
	// Sort is the key to sort the results.
	Sort string
	// Size of the page
	Size int `json:"size"`
	// Offset of the result set
	Offset int `json:"offset"`
}

type ListCharactersResponse struct {
	Characters []entities.Character `json:"characters"`
}
