package main

import (
	"log"

	"github.com/AranGarcia/droop/initiatives/internal/adapters/secondary/redis"
	"github.com/AranGarcia/droop/shared/redistools"
)

func buildRedisConfig() redistools.Config {
	return redistools.Config{
		Addr:     redisHost,
		Password: redisPassword,
	}
}

func initRepository() *redis.TurnRepository {
	config := buildRedisConfig()
	client, err := config.InitializeClient()
	if err != nil {
		log.Fatalf("failed to initialize redis client: %v", err)
	}
	repo := redis.NewTurnsRepository(client)
	return repo
}
