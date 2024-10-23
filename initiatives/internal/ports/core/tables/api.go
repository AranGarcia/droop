package table

import (
	"context"

	"github.com/AranGarcia/initiatives/internal/core/entities"
)

// API is a core port for Tables.
type API interface {
	// RetrieveTable returns the current state of a Campaign's Table and all of the containing turns.
	RetrieveTable(context.Context, RetrieveRequest) (*RetrieveResponse, error)
	// ClearTable resets a Table by deleting all of it's contained turns. It doesn't delete the Table.
	ClearTable(context.Context, ClearRequest) (*ClearResponse, error)
	// StartTracking the turns for a campaign.
	StartTracking(context.Context, StartTrackingRequest) (*StartTrackingResponse, error)
	// StopTracking the turns for a campaign.
	StopTracking(context.Context, StopTrackingRequest) (*StopTrackingResponse, error)
}

type RetrieveRequest struct {
	CampaignID string
}

type RetrieveResponse struct {
	Table entities.Table
}

type ClearRequest struct {
	CampaignID string
}
type ClearResponse struct{}

type StartTrackingRequest struct {
	CampaignID string
}
type StartTrackingResponse struct{}

type StopTrackingRequest struct {
	CampaignID string
}
type StopTrackingResponse struct{}
