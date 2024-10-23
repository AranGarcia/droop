package turns

import (
	"context"

	"github.com/AranGarcia/initiatives/internal/core/entities"
)

type Turns interface {
	Upsert(context.Context, string, entities.Turn) error
	Clear(context.Context, string) error
}
