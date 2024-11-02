package redis

import (
	"context"
	"fmt"
	"time"

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
	t.redisClient.Expire(ctx, campaignID, 4*time.Hour)
	return res.Err()
}

func (t TurnRepository) Clear(ctx context.Context, campaignID string) error {
	res := t.redisClient.Del(ctx, campaignID)
	return res.Err()
}
