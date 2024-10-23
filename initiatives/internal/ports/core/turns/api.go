package turns

import (
	"context"

	"github.com/AranGarcia/initiatives/internal/core/entities"
)

// API is a core port for Turns.
type API interface {
	// Register inserts a Turn into a campaign's tracking registry. If it exists, it overwrites it.
	// It returns the resulting Table after the insert.
	Register(context.Context, RegisterRequest) (*RegisterResponse, error)
	// ClearAll resets a Table by deleting all of it's contained turns. It doesn't delete the Table.
	ClearAll(context.Context, ClearAllRequest) (*ClearAllResponse, error)
}

type RegisterRequest struct {
	CampaignID string
	Turn       entities.Turn
}

type RegisterResponse struct {
	Table entities.Table
}

type ClearAllRequest struct {
	CampaignID string
}
type ClearAllResponse struct{}
