package mongo

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
	"github.com/AranGarcia/droop/dnd/internal/ports/repositories/campaign"
)

type Campaign struct{}

func (d Campaign) CreateCampaign(_ context.Context, _ entities.Campaign) (*entities.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (d Campaign) RetrieveCampaign(_ context.Context, _ string) (*entities.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (d Campaign) UpdateCampaign(_ context.Context, _ campaign.UpdateFields) (*entities.Campaign, error) {
	panic("not implemented") // TODO: Implement
}

func (d Campaign) DeleteCampaign(_ context.Context, _ string) error {
	panic("not implemented") // TODO: Implement
}
