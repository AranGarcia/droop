// Package characters is an external port for the Characters service.
package characters

import (
	"context"

	"github.con/AranGarcia/droop/dnd/core/entities"
)

type Characters interface {
	// Retrieve a Character from the service.
	Retrieve(context.Context, RetrieveCharacterRequest) (*RetrieveCharacterRequest, error)
}

// RetrieveCharacterRequest for Characters.Retrieve.
type RetrieveCharacterRequest struct {
	// ID of the character to retrieve
	ID string `json:"id"`
}

// RetrieveCharacterResponse from Characters.Retrieve.
type RetrieveCharacterResponse struct {
	// Character that was requested.
	Character entities.Character `json:"character"`
}
