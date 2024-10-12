// Package campaign is an outgoing port to access a repository for Campaigns.
package campaign

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
)

// Repository is an outbound port to access the repository for Campaigns.
type Repository interface {
	CreateCampaign(context.Context, entities.Campaign) (*entities.Campaign, error)
	RetrieveCampaign(context.Context, string) (*entities.Campaign, error)
	UpdateCampaign(context.Context, CampaignFields) (*entities.Campaign, error)
	DeleteCampaign(context.Context, string) error
}

type CampaignFields struct{}
