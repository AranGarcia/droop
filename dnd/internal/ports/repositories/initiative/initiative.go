package initiative

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
)

type Repository interface {
	Create(context.Context, entities.Campaign) (*entities.InitiativeTable, error)
	Retrieve(context.Context) (*entities.InitiativeTable, error)
	RetrieveByCampaignID(context.Context) (*entities.InitiativeTable, error)
	AddTurn(context.Context, string) error
	Delete(context.Context, string) error
}
