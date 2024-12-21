package main

import (
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/AranGarcia/droop/dnd/internal/adapters/secondary/characters"
	"github.com/AranGarcia/droop/dnd/internal/core/services"

	adapterkafka "github.com/AranGarcia/droop/dnd/internal/adapters/secondary/kafka"
)

func initDNDAPI() (*services.DND, error) {
	charactersConfig := characters.Config{
		Addr: charactersServiceAddr,
	}
	charactersService, err := characters.NewClient(charactersConfig)
	if err != nil {
		return nil, fmt.Errorf("character service client failure; %v", err)
	}

	w := initKafkaWriter()
	producer := adapterkafka.NewProducer(w)

	dndService := services.NewDNDService(charactersService, producer)
	return &dndService, nil
}

func initKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(kafkaBrokerAddr),
		Topic: adapterkafka.Topic,
	}
}
