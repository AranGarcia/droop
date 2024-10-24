package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
)

type TablesRepository struct {
	client redis.Client
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
