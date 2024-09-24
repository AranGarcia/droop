package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/dnd/internal/adapters/primary/grpc"
)

var (
	addr string
	// charactersServiceAddr is the host where Character Service is running on.
	charactersServiceAddr string
)

func init() {
	flag.StringVar(&addr, "addr", "8010", "the server address (host:port)")
	flag.Parse()
}

func main() {
	log.Println("Running the server on", addr)
	dndAPI, err := initDNDAPI()
	if err != nil {
		log.Fatalf("failed to initialize api; %v", err)
	}

	c := grpc.Config{
		DND: dndAPI,
	}
	server := grpc.NewServer(c)
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run failure; %v", err)
	}
}
