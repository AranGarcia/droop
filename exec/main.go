package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/internal"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "the server address (host and port)")
	flag.Parse()
}

func main() {
	log.Println("Running HTTP server...")
	httpServer := internal.NewServer(addr)

	if err := httpServer.Run(); err != nil {
		log.Fatalf("Failed to run; %v", err)
	}
}
