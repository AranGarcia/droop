package turns

import (
	"context"

	"github.com/AranGarcia/droop/initiatives/internal/ports/out/repositories/turns"

	core "github.com/AranGarcia/droop/initiatives/internal/ports/core/turns"
)

type Dependencies struct {
	Repository turns.Repository
}

type Service struct {
	repo turns.Repository
}

func NewService(deps Dependencies) *Service {
	return &Service{
		repo: deps.Repository,
	}
}

func (s Service) ListTurns(ctx context.Context, request core.ListTurnsRequest) (*core.ListTurnsResponse, error) {
	response := &core.ListTurnsResponse{}
	return response, nil
}

// Register inserts a Turn into a campaign's tracking registry. If it exists, it overwrites it.
// It returns the resulting Table after the insert.
func (s Service) Register(ctx context.Context, request core.RegisterRequest) (*core.RegisterResponse, error) {
	if err := s.repo.Upsert(ctx, request.CampaignID, request.Turn); err != nil {
		return nil, err // TODO: handle errors
	}
	response := &core.RegisterResponse{}
	return response, nil
}

// ClearTable resets a Table by deleting all of it's contained turns. It doesn't delete the Table.
func (s Service) ClearAll(ctx context.Context, request core.ClearAllRequest) (*core.ClearAllResponse, error) {
	if err := s.repo.Clear(ctx, request.CampaignID); err != nil {
		return nil, err // TODO: handle errors
	}
	response := &core.ClearAllResponse{}
	return response, nil
}
