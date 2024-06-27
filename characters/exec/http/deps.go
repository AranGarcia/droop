package main

import (
	"log"
	"time"

	"github.com/AranGarcia/droop/characters/internal/adapters/secondary/mongo"
	"github.com/AranGarcia/droop/characters/internal/core/services"
	"github.com/AranGarcia/droop/characters/internal/ports/api"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

// buildMongoConfig is built from the configuration flags.
func buildMongoConfig() mongo.Config {
	return mongo.Config{
		User:     mongoUser,
		Password: mongoPassword,
		Host:     mongoHost,
		Port:     mongoPort,
		Database: mongoDatabase,
		Timeout:  time.Second * 5,
	}
}

func buildRepository(mongoConfig mongo.Config) repositories.Characters {
	repo, err := mongo.NewCharacterRepository(mongoConfig)
	if err != nil {
		log.Fatalf("failed to initialize repo; %v", err)
	}
	return repo
}

func buildService(repo repositories.Characters) api.Characters {
	return &services.Characters{Repository: repo}
}
