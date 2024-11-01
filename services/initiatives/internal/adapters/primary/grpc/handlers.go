package grpc

import (
	"context"

	"github.com/AranGarcia/droop/initiatives/internal/ports/core/tables"

	initiativespb "github.com/AranGarcia/droop/proto/gen/initiatives"
)

func (t Server) StartTracking(ctx context.Context, request *initiativespb.StartTrackingRequest) (*initiativespb.StartTrackingResponse, error) {
	apiRequest := tables.StartTrackingRequest{CampaignID: request.CampaignId}
	_, err := t.tablesAPI.StartTracking(ctx, apiRequest)
	if err != nil {
		// TODO: handle core errors
		return nil, err
	}
	response := &initiativespb.StartTrackingResponse{}
	return response, nil
}

func (t Server) StopTracking(ctx context.Context, request *initiativespb.StopTrackingRequest) (*initiativespb.StopTrackingResponse, error) {
	apiRequest := tables.StopTrackingRequest{CampaignID: request.CampaignId}
	_, err := t.tablesAPI.StopTracking(ctx, apiRequest)
	if err != nil {
		// TODO: handle core errors
		return nil, err
	}
	response := &initiativespb.StopTrackingResponse{}
	return response, nil
}
