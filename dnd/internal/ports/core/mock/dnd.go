package mock

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/ports/core/dnd"
)

// DND is a mock implementation of the core.
type DND struct {
	RollInitiativeFunc func() (*dnd.RollInitiativeResponse, error)
}

// RollInitiative just calls DND.RollInitiativeFunc
func (d DND) RollInitiative(_ context.Context, _ dnd.RollInitiativeRequest) (*dnd.RollInitiativeResponse, error) {
	return d.RollInitiativeFunc()
}
