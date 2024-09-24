// Package characters is an external port for the Characters service.
package characters

import (
	"context"

	"github.con/AranGarcia/droop/dnd/internal/core/entities"
)

// Port to the external service Characters.
type Port interface {
	// Retrieve a Character from the service.
	Retrieve(context.Context, string) (*entities.Character, error)
}
