package service

import (
	"context"
	"fmt"

	"github.com/AranGarcia/droop/campaigns/internal/core/api"
	"github.com/AranGarcia/droop/campaigns/internal/core/entities"
)

func (c Campaigns) Create(ctx context.Context, _ api.CreateRequest) (*api.CreateResponse, error) {
	entity := entities.Campaign{}
	if err := c.repository.Create(ctx, entity); err != nil {
		return nil, fmt.Errorf("failed to create; %v", err)
	}
	response := &api.CreateResponse{}
	return response, nil
}
