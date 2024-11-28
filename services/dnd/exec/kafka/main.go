package main

import (
	"log"

	"github.com/AranGarcia/droop/dnd/internal/adapters/primary/kafka"
)

func main() {
	_, err := kafka.NewConsumer()
	if err != nil {
		log.Fatalf("failed to start consumer; %v", err)
	}
	log.Println("Success")
}
