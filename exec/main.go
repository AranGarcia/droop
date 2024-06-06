package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/internal"
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
)

func init() {
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "the server address (host and port)")
	flag.StringVar(&mongoUser, "mongo_user", "droopadmin", "MongoDB user name")
	flag.StringVar(&mongoPassword, "mongo_password", "droopadmin", "MongoDB user password")
	flag.StringVar(&mongoHost, "mongo_host", "localhost", "MongoDB hostname")
	flag.IntVar(&mongoPort, "mongo_port", 27017, "MongoDB port")

	flag.Parse()
}

func main() {
	log.Println("Initializing repository...")
	repository, err := internal.NewCharacterRepository(mongoUser, mongoPassword, mongoHost, mongoPort)
	if err != nil {
		log.Fatal("mongo client initialization failed; ", err)
	}
	handler := internal.NewHandler(repository)
	log.Println("running the server...")
	httpServer := internal.NewServer(addr, handler)
	log.Println("Server is running on", addr)
	if err := httpServer.Run(); err != nil {
		log.Fatalf("Failed to run; %v", err)
	}
}
