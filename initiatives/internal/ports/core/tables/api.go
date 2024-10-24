package tables

import (
	"context"
)

// API is a core port for Tables.
type API interface {
	// StartTracking the turns for a campaign.
	StartTracking(context.Context, StartTrackingRequest) (*StartTrackingResponse, error)
	// StopTracking the turns for a campaign.
	StopTracking(context.Context, StopTrackingRequest) (*StopTrackingResponse, error)
}

type StartTrackingRequest struct {
	CampaignID string
}
type StartTrackingResponse struct{}

type StopTrackingRequest struct {
	CampaignID string
}
type StopTrackingResponse struct{}
