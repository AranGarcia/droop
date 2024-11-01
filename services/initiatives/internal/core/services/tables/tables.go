package tables

import (
	"context"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
	"github.com/AranGarcia/droop/initiatives/internal/ports/out/repositories/tables"

	coreerrors "github.com/AranGarcia/droop/initiatives/internal/ports/core"
	tablescore "github.com/AranGarcia/droop/initiatives/internal/ports/core/tables"
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
func (s Service) StartTracking(ctx context.Context, request tablescore.StartTrackingRequest) (*tablescore.StartTrackingResponse, error) {
	if request.CampaignID == "" {
		return nil, coreerrors.ErrInvalidID
	}

	table := entities.Table{
		CampaignID: request.CampaignID,
	}
	if err := s.repository.Create(ctx, table); err != nil {
		return nil, err // TODO: handle errors
	}

	response := &tablescore.StartTrackingResponse{}
	return response, nil
}

// StopTracking the turns for a campaign.
func (s Service) StopTracking(ctx context.Context, request tablescore.StopTrackingRequest) (*tablescore.StopTrackingResponse, error) {
	if err := s.repository.Delete(ctx, request.CampaignID); err != nil {
		return nil, err // TODO: handle errors
	}

	response := &tablescore.StopTrackingResponse{}
	return response, nil
}
