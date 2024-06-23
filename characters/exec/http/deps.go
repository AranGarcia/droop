package main

import (
	"time"

	"github.com/AranGarcia/droop/characters/internal"
	"github.com/AranGarcia/droop/characters/internal/adapters/secondary/mongo"
	"github.com/AranGarcia/droop/characters/internal/core/services"
	"github.com/AranGarcia/droop/characters/internal/ports/api"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
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

func buildRepository(mongoConfig mongo.Config) repositories.Characters {
	return mongo.CharacterRepository{}
}

func buildService(repo repositories.Characters) api.Characters {
	return &services.Characters{Repository: repo}
}
