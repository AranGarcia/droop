package main

import (
	"log"
	"time"

	"github.com/AranGarcia/droop/characters/internal/adapters/secondary/mongo"
	"github.com/AranGarcia/droop/characters/internal/core/services"
	"github.com/AranGarcia/droop/characters/internal/ports/api"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
	"github.com/AranGarcia/droop/shared/mongotools"
)

// buildMongoConfig is built from the configuration flags.
func buildMongoConfig() mongotools.Config {
	return mongotools.Config{
		User:     mongoUser,
		Password: mongoPassword,
		Host:     mongoHost,
		Port:     mongoPort,
		Database: mongoDatabase,
		Timeout:  time.Second * 5,
	}
}

func buildRepository(mongoConfig mongotools.Config) repositories.Characters {
	repo, err := mongo.NewCharacterRepository(mongoConfig)
	if err != nil {
		log.Fatalf("failed to initialize repo; %v", err)
	}
	return repo
}

func buildService(repo repositories.Characters) api.Characters {
	deps := services.Dependencies{Repository: repo}
	return services.NewCharacters(deps)
}
