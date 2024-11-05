package turns

import (
	"context"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
)

// Repository for the the turns of a campaign.
type Repository interface {
	// List the turns for a given campaign.
	List(context.Context, string) ([]entities.Turn, error)
	// Upsert a turn into the campaign's initiative table.
	Upsert(context.Context, string, entities.Turn) error
	// Clear all of the turns for a campaign.
	Clear(context.Context, string) error
}
