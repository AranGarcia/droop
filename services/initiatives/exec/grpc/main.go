package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/initiatives/internal/adapters/primary/grpc"
	"github.com/AranGarcia/droop/initiatives/internal/core/services/tables"
)

var (
	// addr is the address (HOST:PORT) where the server will run.
	addr string

	// redisHost is the hostname of the Redis
	redisHost     string
	redisPassword string
)

func init() {
	flag.StringVar(&addr, "addr", ":8070", "the server address (host and port)")
	flag.StringVar(&redisHost, "redis_host", ":6379", "host and port for the Redis database")
	flag.StringVar(&redisPassword, "redis_password", "password", "password for the redis user")
	flag.Parse()
}

func main() {
	repo := buildRepository()
	deps := tables.Dependencies{
		Repository: repo,
	}
	tableService := tables.NewService(deps)
	server := grpc.NewServer(addr, tableService)
	log.Println("Running the server on", addr)
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run failure; %v", err)
	}
}
