package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/characters/internal/adapters/primary/grpc"
)

var (
	// addr is the address (HOST:PORT) where the server will run.
	addr string
	// mongoUser is the username that will connect to the MongoDB.
	mongoUser string
	// mongoPassword is the password of the user.
	mongoPassword string
	// mongoHost is the hostname of the MongoDB.
	mongoHost string
	// mongoPort of the MongoDB.
	mongoPort int
	// mongoDatabase is the name of the database where collections are stored.
	mongoDatabase string
)

func init() {
	flag.StringVar(&addr, "addr", ":8090", "the server address (host and port)")
	flag.StringVar(&mongoUser, "mongo_user", "droopadmin", "MongoDB user name")
	flag.StringVar(&mongoPassword, "mongo_password", "droopadmin", "MongoDB user password")
	flag.StringVar(&mongoHost, "mongo_host", "localhost", "MongoDB hostname")
	flag.IntVar(&mongoPort, "mongo_port", 27017, "MongoDB port")
	flag.StringVar(&mongoDatabase, "mongo_database", "droop", "MongoDB database name")
	flag.Parse()
}

func main() {
	log.Printf("Connection to repository in host %s...\n", mongoHost)
	mongoConfig := buildMongoConfig()
	repo := buildRepository(mongoConfig)
	service := buildService(repo)
	server := grpc.NewServer(addr, service)

	log.Println("Running the server on", addr)
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run failure; %v", err)
	}
}
