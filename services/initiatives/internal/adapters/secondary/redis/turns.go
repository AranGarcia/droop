package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
)

type TurnRepository struct {
	redisClient *redis.Client
}

func NewTurnsRepository(redisClient *redis.Client) *TurnRepository {
	repo := TurnRepository{redisClient: redisClient}
	return &repo
}

func (t TurnRepository) Upsert(_ context.Context, _ string, _ entities.Turn) error {
	panic("not implemented") // TODO: Implement
}

func (t TurnRepository) Clear(_ context.Context, _ string) error {
	panic("not implemented") // TODO: Implement
}
