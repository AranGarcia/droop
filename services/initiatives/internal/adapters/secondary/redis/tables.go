package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
	"github.com/AranGarcia/droop/shared/redistools"
)

// TablesRepository ...
// Deprecated: Use only the TurnsRepository.
type TablesRepository struct {
	client redis.Client
}

// NewTablesRepository ...
// Deprecated: Use only the TurnsRepository.
func NewTablesRepository(config redistools.Config) (*TablesRepository, error) {
	client, err := config.InitializeClient()
	if err != nil {
		return nil, err
	}
	repo := &TablesRepository{client: *client}
	return repo, nil
}

func (t TablesRepository) Create(ctx context.Context, table entities.Table) error {
	res := t.client.SAdd(ctx, "initiatives:campaigns", table.CampaignID)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func (t TablesRepository) Retrieve(_ context.Context, _ string) (*entities.Table, error) {
	panic("not implemented") // TODO: Implement
}

func (t TablesRepository) Delete(_ context.Context, _ string) error {
	panic("not implemented") // TODO: Implement
}
