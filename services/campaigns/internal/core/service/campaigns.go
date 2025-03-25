package service

import (
	"context"

	"github.com/AranGarcia/droop/campaigns/internal/core/api"
	"github.com/AranGarcia/droop/campaigns/internal/ports/campaigns"
	"github.com/AranGarcia/droop/campaigns/internal/ports/characters"
)

type Campaigns struct {
	repository     campaigns.Port
	charactersPort characters.Port
}

func (c Campaigns) Retrieve(_ context.Context, _ api.RetrieveRequest) (*api.RetrieveResponse, error) {
	panic("not implemented")
}

func (c Campaigns) Update(_ context.Context, _ api.UpdateRequest) (*api.UpdateResponse, error) {
	panic("not implemented")
}

func (c Campaigns) Delete(_ context.Context, _ api.DeleteRequest) (*api.DeleteResponse, error) {
	panic("not implemented")
}

func (c Campaigns) List(_ context.Context, _ api.ListRequest) (*api.ListResponse, error) {
	panic("not implemented")
}
