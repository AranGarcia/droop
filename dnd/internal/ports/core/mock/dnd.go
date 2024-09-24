package mock

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/ports/core"
)

// DND is a mock implementation of the core.
type DND struct {
	RollInitiativeFunc func() (*core.RollInitiativeResponse, error)
}

// RollInitiative just calls DND.RollInitiativeFunc
func (d DND) RollInitiative(_ context.Context, _ core.RollInitiativeRequest) (*core.RollInitiativeResponse, error) {
	return d.RollInitiativeFunc()
}
