package kafka

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/ports/events"
)

type Producer struct{}

func NewProducer() Producer {
	return Producer{}
}

// RollInitiativeSuccess produces the event of a successful execution of RollInitiative.
func (p Producer) RollInitiativeSuccess(_ context.Context, _ events.RollInitiativeSuccessMessage) error {
	panic("not implemented") // TODO: Implement
}
