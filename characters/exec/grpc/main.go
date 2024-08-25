package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/characters/internal/adapters/primary/grpc"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8090", "the server address (host:port)")
	flag.Parse()
}

func main() {
	log.Println("Running the server on", addr)
	server := grpc.NewServer(addr)
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run failure; %v", err)
	}
}
