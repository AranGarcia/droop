package mock

import (
	"context"
	"time"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
	"github.com/AranGarcia/droop/dnd/internal/ports/core"
	"github.com/AranGarcia/droop/dnd/internal/ports/core/campaign"
)

type Campaign struct {
	Entities map[string]entities.Campaign
}

func NewMockCampaignService() *Campaign {
	return &Campaign{
		Entities: make(map[string]entities.Campaign),
	}
}

// Create a Campaign.
func (c *Campaign) Create(_ context.Context, request campaign.CreateRequest) (*campaign.CreateResponse, error) {
	id := time.Now().String()
	entity := entities.Campaign{
		ID:   id,
		Name: request.Name,
	}
	c.Entities[id] = entity
	response := &campaign.CreateResponse{
		Campaign: entity,
	}
	return response, nil
}

// Retrieve a Campaign.§
func (c *Campaign) Retrieve(_ context.Context, request campaign.RetrieveRequest) (*campaign.RetrieveResponse, error) {
	entity, ok := c.Entities[request.ID]
	if !ok {
		return nil, core.ErrNotFound
	}
	response := &campaign.RetrieveResponse{Campaign: entity}
	return response, nil
}
