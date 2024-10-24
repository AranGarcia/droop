package tables

import (
	"context"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
)

type Repository interface {
	Create(context.Context, entities.Table) error
	Retrieve(context.Context, string) (*entities.Table, error)
	Delete(context.Context, string) error
}
