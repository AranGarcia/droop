package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/initiatives/internal/adapters/primary/grpc"
	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
	"github.com/AranGarcia/droop/initiatives/internal/core/services/tables"
	"github.com/AranGarcia/droop/initiatives/internal/ports/out/repositories/mocks"
)

var (
	// addr is the address (HOST:PORT) where the server will run.
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8070", "the server address (host and port)")
	flag.Parse()
}

func main() {
	deps := tables.Dependencies{
		// TODO: change for redis adapter
		Repository: &mocks.TableRepository{
			Tables: make(map[string]*entities.Table),
		},
	}
	tableService := tables.NewService(deps)
	server := grpc.NewServer(addr, tableService)
	log.Println("Running the server on", addr)
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run failure; %v", err)
	}
}
