package tables

import (
	"context"

	"github.com/AranGarcia/initiatives/internal/core/entities"
)

// API is a core port for Tables.
type API interface {
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

type StartTrackingRequest struct {
	CampaignID string
}
type StartTrackingResponse struct{}

type StopTrackingRequest struct {
	CampaignID string
}
type StopTrackingResponse struct{}
