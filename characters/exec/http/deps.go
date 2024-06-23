package main

import (
	"log"
	"time"

	"github.com/AranGarcia/droop/characters/internal"
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

type repositories struct {
	characters internal.CharacterRepository
}

func buildRepositories(mongoConfig internal.MongoConfig) repositories {
	characters, err := internal.NewCharacterRepository(mongoConfig)
	if err != nil {
		log.Fatal("mongo client initialization failed; ", err)
	}
	return repositories{
		characters: *characters,
	}
}
