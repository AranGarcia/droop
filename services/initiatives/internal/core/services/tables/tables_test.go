package tables

import (
	"context"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"
	core "github.com/AranGarcia/droop/initiatives/internal/ports/core/tables"
	"github.com/AranGarcia/droop/initiatives/internal/ports/out/repositories/mocks"
)

func TestService_StartTracking(t *testing.T) {
	s := Service{
		repository: &mocks.TableRepository{
			Tables: map[string]*entities.Table{
				"existing-campaign-id": {CampaignID: "existing-campaign-id"},
			},
		},
	}

	tests := []struct {
		name    string
		request core.StartTrackingRequest
		want    *core.StartTrackingResponse
		wantErr bool
	}{
		{
			name:    "create error",
			request: core.StartTrackingRequest{CampaignID: "existing-campaign-id"},
			wantErr: true,
		},
		{
			name:    "start tracking success",
			request: core.StartTrackingRequest{CampaignID: "new-campaign-id"},
			want:    &core.StartTrackingResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()
			got, err := s.StartTracking(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.StartTracking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.StartTracking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_StopTracking(t *testing.T) {
	s := Service{
		repository: &mocks.TableRepository{
			Tables: map[string]*entities.Table{
				"existing-campaign-id": {CampaignID: "existing-campaign-id"},
			},
		},
	}

	tests := []struct {
		name    string
		request core.StopTrackingRequest
		want    *core.StopTrackingResponse
		wantErr bool
	}{
		{
			name:    "delete error",
			request: core.StopTrackingRequest{CampaignID: "non-existing-campaign-id"},
			wantErr: true,
		},
		{
			name:    "stop tracking success",
			request: core.StopTrackingRequest{CampaignID: "existing-campaign-id"},
			want:    &core.StopTrackingResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := s.StopTracking(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.StopTracking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.StopTracking() = %v, want %v", got, tt.want)
			}
		})
	}
}
