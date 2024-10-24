package tables

import (
	"context"
	"reflect"
	"testing"

	"github.com/AranGarcia/initiatives/internal/core/entities"
	"github.com/AranGarcia/initiatives/internal/ports/out/repositories/mocks"

	core "github.com/AranGarcia/initiatives/internal/ports/core/tables"
)

func TestService_StartTracking(t *testing.T) {
	s := Service{
		repository: &mocks.TableRepository{
			Tables: map[string]*entities.Table{
				"existing-campaign-id": {CampaignID: "existing-campaign-id"},
			},
		},
	}

	type args struct {
		ctx     context.Context
		request core.StartTrackingRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *core.StartTrackingResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.StartTracking(tt.args.ctx, tt.args.request)
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
