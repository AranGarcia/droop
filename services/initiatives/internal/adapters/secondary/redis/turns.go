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
	args := redis.ZRangeArgs{
		Key:   campaignID,
		Start: 0,
		Stop:  -1,
		Rev:   true,
	}
	res := t.redisClient.ZRangeArgsWithScores(ctx, args)
	results, err := res.Result()
	if err != nil {
		return nil, fmt.Errorf("failed to extract results; %v", err)
	}
	turns := make([]entities.Turn, len(results))
	for i, z := range results {
		turns[i], err = ZMemberToTurn(z)
		if err != nil {
			return nil, err
		}
	}
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
