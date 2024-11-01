package main

import (
	"log"

	"github.com/AranGarcia/droop/initiatives/internal/adapters/secondary/redis"
	"github.com/AranGarcia/shared/redistools"
)

func buildRedisConfig() redistools.Config {
	return redistools.Config{
		Addr:     redisHost,
		Password: redisPassword,
	}
}

func buildRepository() *redis.TablesRepository {
	config := buildRedisConfig()
	repo, err := redis.NewTablesRepository(config)
	if err != nil {
		log.Fatalf("failed to initialize tables repository; %v", err)
	}
	return repo
}
