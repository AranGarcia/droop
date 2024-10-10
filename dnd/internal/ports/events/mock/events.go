package mock

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/ports/events"
)

// Producer is a mock implementation of the event port.
type Producer struct {
	eventsProduced int
}

// RollInitiativeSuccess produces the event of a successful execution of RollInitiative.
func (p *Producer) RollInitiativeSuccess(_ context.Context, _ events.RollInitiativeSuccessMessage) error {
	p.eventsProduced += 1
	return nil
}

func (p *Producer) AmountOfMessagesProduced() int { return p.eventsProduced }
