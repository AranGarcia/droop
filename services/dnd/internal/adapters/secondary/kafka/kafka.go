package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"

	"github.com/AranGarcia/droop/dnd/internal/ports/events"
)

const Topic = ""

type Producer struct {
	w *kafka.Writer
}

func NewProducer(w *kafka.Writer) *Producer {
	return &Producer{w: w}
}

// RollInitiativeSuccess produces the event of a successful execution of RollInitiative.
func (p *Producer) RollInitiativeSuccess(ctx context.Context, m events.RollInitiativeSuccessMessage) error {
	message := kafka.Message{
		Key:   []byte("test"),
		Value: []byte("initiative-rolled"),
	}
	return p.w.WriteMessages(ctx, message)
}
