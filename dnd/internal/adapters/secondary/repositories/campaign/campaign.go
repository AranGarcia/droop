package campaign

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/ports/repositories/campaign"
)

type Database struct{}

func (d Database) CreateCampaign(_ context.Context, _ campaign.CreateCampaignRequest) (*campaign.CreateCampaignResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (d Database) RetrieveCampaign(_ context.Context, _ campaign.RetrieveCampaignRequest) (*campaign.RetrieveCampaignResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (d Database) UpdateCampaign(_ context.Context, _ campaign.UpdateCampaignRequest) (*campaign.UpdateCampaignResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (d Database) DeleteCampaign(_ context.Context, _ campaign.DeleteCampaignRequest) (*campaign.DeleteCampaignResponse, error) {
	panic("not implemented") // TODO: Implement
}
