package main

import (
	"log"
	"time"

	"github.com/AranGarcia/droop/internal"
)

// buildMongoConfig is built from the configuration flags.
func buildMongoConfig() internal.MongoConfig {
	return internal.MongoConfig{
		User:     mongoUser,
		Password: mongoPassword,
		Host:     mongoHost,
		Port:     mongoPort,
		Database: mongoDatabase,
		Timeout:  time.Second * 5,
	}
}

// buildCharacterRepository initializes the character
func buildCharacterRepository(m internal.MongoConfig) *internal.CharacterRepository {
	repository, err := internal.NewCharacterRepository(m)
	if err != nil {
		log.Fatal("mongo client initialization failed; ", err)
	}
	return repository
}
