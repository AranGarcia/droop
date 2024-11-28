package kafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	conn *kafka.Conn
}

func NewConsumer() (*Consumer, error) {
	c := &Consumer{}
	var err error
	c.conn, err = kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		return nil, fmt.Errorf("failed to dial; %v", err)
	}
	return c, nil
}
