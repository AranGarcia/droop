package tables

import (
	"context"

	// XXX: Not a fan of this import naming scheme. Is there a better way to do this?

	"github.com/AranGarcia/initiatives/internal/core/entities"
	"github.com/AranGarcia/initiatives/internal/ports/out/repositories/tables"

	core "github.com/AranGarcia/initiatives/internal/ports/core/tables"
)

type Service struct {
	repository tables.Repository
}

// RetrieveTable returns the current state of a Campaign's Table and all of the containing turns.
func (s Service) RetrieveTable(ctx context.Context, request core.RetrieveRequest) (*core.RetrieveResponse, error) {
	table, err := s.repository.Retrieve(ctx, request.CampaignID)
	if err != nil {
		return nil, err // TODO: handle errors
	}
	response := &core.RetrieveResponse{
		Table: *table,
	}
	return response, nil
}

// ClearTable resets a Table by deleting all of it's contained turns. It doesn't delete the Table.
func (s Service) ClearTable(_ context.Context, _ core.ClearRequest) (*core.ClearResponse, error) {
	panic("not implemented") // TODO: Implement
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
