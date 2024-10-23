package turns

import (
	"context"

	"github.com/AranGarcia/initiatives/internal/ports/out/repositories/tables"
	"github.com/AranGarcia/initiatives/internal/ports/out/repositories/turns"

	core "github.com/AranGarcia/initiatives/internal/ports/core/turns"
)

type Service struct {
	// Repositories
	turnRepo  turns.Repository
	tableRepo tables.Repository
}

func (s Service) ListTurns(ctx context.Context) (*core.ListTurnsResponse, error) {
	response := &core.ListTurnsResponse{}
	return response, nil
}

// Register inserts a Turn into a campaign's tracking registry. If it exists, it overwrites it.
// It returns the resulting Table after the insert.
func (s Service) Register(ctx context.Context, request core.RegisterRequest) (*core.RegisterResponse, error) {
	// If table exists, it means its tracking
	_, err := s.tableRepo.Retrieve(ctx, request.CampaignID)
	if err != nil {
		// TODO: handle when error is caused by NOT FOUND.
		return nil, err
	}

	if err := s.turnRepo.Upsert(ctx, request.CampaignID, request.Turn); err != nil {
		return nil, err // TODO: handle errors
	}

	response := &core.RegisterResponse{}
	return response, nil
}

// ClearTable resets a Table by deleting all of it's contained turns. It doesn't delete the Table.
func (s Service) ClearAll(ctx context.Context, request core.ClearAllRequest) (*core.ClearAllResponse, error) {
	if err := s.turnRepo.Clear(ctx, request.CampaignID); err != nil {
		return nil, err // TODO: handle errors
	}
	response := &core.ClearAllResponse{}
	return response, nil
}
