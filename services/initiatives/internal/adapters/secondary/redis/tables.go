package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
	"github.com/AranGarcia/shared/redistools"
)

type TablesRepository struct {
	client redis.Client
}

func NewTablesRepository(config redistools.Config) (*TablesRepository, error) {
	client, err := config.InitializeClient()
	if err != nil {
		return nil, err
	}
	repo := &TablesRepository{client: *client}
	return repo, nil
}

func (t TablesRepository) Create(ctx context.Context, _ entities.Table) error {
	_ = t.client.SAdd(ctx, "", nil)
	return nil
}

func (t TablesRepository) Retrieve(_ context.Context, _ string) (*entities.Table, error) {
	panic("not implemented") // TODO: Implement
}

func (t TablesRepository) Delete(_ context.Context, _ string) error {
	panic("not implemented") // TODO: Implement
}
