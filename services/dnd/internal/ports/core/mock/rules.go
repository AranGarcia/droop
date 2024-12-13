package mock

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/ports/core/rules"
)

// Rules is a mock implementation of the core.
type Rules struct {
	RollInitiativeFunc func() (*rules.RollInitiativeResponse, error)
}

// RollInitiative just calls DND.RollInitiativeFunc
func (d Rules) RollInitiative(_ context.Context, _ rules.RollInitiativeRequest) (*rules.RollInitiativeResponse, error) {
	return d.RollInitiativeFunc()
}
