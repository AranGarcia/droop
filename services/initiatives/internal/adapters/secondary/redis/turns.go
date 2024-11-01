package redis

import (
	"context"
	"fmt"

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

func (t TurnRepository) Upsert(ctx context.Context, campaignID string, turn entities.Turn) error {
	z := redis.Z{
		Score:  float64(turn.Result),
		Member: fmt.Sprint(turn.CharacterID, ":", turn.CharacterName),
	}
	res := t.redisClient.ZAdd(ctx, campaignID, z)
	return res.Err()
}

func (t TurnRepository) Clear(_ context.Context, _ string) error {
	panic("not implemented") // TODO: Implement
}
