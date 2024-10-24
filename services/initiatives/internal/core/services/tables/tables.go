package tables

import (
	"context"

	// XXX: Not a fan of this import naming scheme. Is there a better way to do this?

	"github.com/AranGarcia/initiatives/internal/core/entities"
	"github.com/AranGarcia/initiatives/internal/ports/out/repositories/tables"

	core "github.com/AranGarcia/initiatives/internal/ports/core/tables"
)

type Dependencies struct {
	Repository tables.Repository
}

type Service struct {
	repository tables.Repository
}

func NewService(deps Dependencies) Service {
	return Service{
		repository: deps.Repository,
	}
}

// StartTracking the turns for a campaign.
func (s Service) StartTracking(ctx context.Context, request core.StartTrackingRequest) (*core.StartTrackingResponse, error) {
	table := entities.Table{
		CampaignID: request.CampaignID,
	}
	if err := s.repository.Create(ctx, table); err != nil {
		return nil, err // TODO: handle errors
	}

	response := &core.StartTrackingResponse{}
	return response, nil
}

// StopTracking the turns for a campaign.
func (s Service) StopTracking(ctx context.Context, request core.StopTrackingRequest) (*core.StopTrackingResponse, error) {
	if err := s.repository.Delete(ctx, request.CampaignID); err != nil {
		return nil, err // TODO: handle errors
	}

	response := &core.StopTrackingResponse{}
	return response, nil
}
