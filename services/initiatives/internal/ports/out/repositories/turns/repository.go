package turns

import (
	"context"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
)

type Repository interface {
	Upsert(context.Context, string, entities.Turn) error
	Clear(context.Context, string) error
}
