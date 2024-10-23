package turns

import (
	"context"

	"github.com/AranGarcia/initiatives/internal/ports/core/turns"

	repository "github.com/AranGarcia/initiatives/internal/ports/out/repositories/turns"
)

type Service struct {
	repo repository.Repository
}

// Register inserts a Turn into a campaign's tracking registry. If it exists, it overwrites it.
// It returns the resulting Table after the insert.
func (s Service) Register(ctx context.Context, request turns.RegisterRequest) (*turns.RegisterResponse, error) {
	if err := s.repo.Upsert(ctx, request.CampaignID, request.Turn); err != nil {
		return nil, err // TODO: handle errors
	}

	response := &turns.RegisterResponse{}
	return response, nil
}

// ClearTable resets a Table by deleting all of it's contained turns. It doesn't delete the Table.
func (s Service) ClearAll(ctx context.Context, request turns.ClearAllRequest) (*turns.ClearAllResponse, error) {
	if err := s.repo.Clear(ctx, request.CampaignID); err != nil {
		return nil, err // TODO: handle errors
	}
	response := &turns.ClearAllResponse{}
	return response, nil
}
