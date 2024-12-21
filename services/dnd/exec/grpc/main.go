package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/dnd/internal/adapters/primary/grpc"
)

var (
	addr                  string
	charactersServiceAddr string
	kafkaBrokerAddr       string
)

func init() {
	flag.StringVar(&addr, "addr", ":8010", "the server address (host:port)")
	flag.StringVar(&charactersServiceAddr, "character_service_addr", ":8090", "host where Character Service is running")
	flag.StringVar(&kafkaBrokerAddr, "kafka_broker_addr", ":9092", "")

	flag.Parse()
}

func main() {
	log.Println("Running the server on", addr)
	dndAPI, err := initDNDAPI()
	if err != nil {
		log.Fatalf("failed to initialize api; %v", err)
	}

	c := grpc.Config{
		Addr: addr,
		DND:  dndAPI,
	}
	server := grpc.NewServer(c)
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run failure; %v", err)
	}
}
