package campaigns

import (
	"context"

	"github.com/AranGarcia/droop/campaigns/internal/core/entities"
)

type Port interface {
	Create(context.Context, entities.Campaign) error
	Retrieve(context.Context, string) (*entities.Campaign, error)
	Update(context.Context, UpdateFields) (*entities.Campaign, error)
	Delete(context.Context, string) error
}

type UpdateFields struct{}
