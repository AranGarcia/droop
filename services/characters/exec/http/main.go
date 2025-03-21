package main

import (
	"flag"
	"log"

	"github.com/AranGarcia/droop/characters/internal/adapters/primary/http"
)

var (
	// addr is the address (HOST:PORT) where the server will run.
	addr string
	// accessControlAllowOrigin is used to set response headers for CORS policies.
	accessControlAllowOrigin string
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
	flag.StringVar(&addr, "addr", "0.0.0.0:8080", "the server address (host and port)")
	flag.StringVar(&accessControlAllowOrigin, "access_control_allow_origin", "*", "will be set in response headers for CORS")
	flag.StringVar(&mongoUser, "mongo_user", "droopadmin", "MongoDB user name")
	flag.StringVar(&mongoPassword, "mongo_password", "droopadmin", "MongoDB user password")
	flag.StringVar(&mongoHost, "mongo_host", "localhost", "MongoDB hostname")
	flag.IntVar(&mongoPort, "mongo_port", 27017, "MongoDB port")
	flag.StringVar(&mongoDatabase, "mongo_database", "droop", "MongoDB database name")
	flag.Parse()
}

func main() {
	log.Println("Initializing repository...")
	mongoConfig := buildMongoConfig()
	repo := buildRepository(mongoConfig)
	service := buildService(repo)
	server := http.NewServer(addr, accessControlAllowOrigin, service)
	log.Println("Running server on", addr)
	if err := server.Run(); err != nil {
		log.Fatalf("server.Run failure; %v", err)
	}
}
