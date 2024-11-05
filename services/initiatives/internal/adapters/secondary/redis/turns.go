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

// List the turns for a given campaign.
func (t TurnRepository) List(ctx context.Context, campaignID string) ([]entities.Turn, error) {
	res := t.redisClient.ZRange(ctx, campaignID, 0, -1)
	if res.Err() != nil {
		return nil, fmt.Errorf("failed to range over set; %v", res.Err())
	}
	results, err := res.Result() // TODO: make sure that this and res.Err() could be two separate errors.
	if err != nil {
		return nil, fmt.Errorf("failed to extract results; %v", err)
	}
	turns := make([]entities.Turn, len(results))
	// TODO: parse result into an entities.Turn
	return turns, nil
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
