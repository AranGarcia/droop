package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/campaigns/internal/core/api"
	"github.com/AranGarcia/droop/campaigns/internal/ports/campaigns"
	"github.com/AranGarcia/droop/campaigns/internal/ports/characters"
)

func TestCampaigns_Create(t *testing.T) {
	type fields struct {
		repository     campaigns.Port
		charactersPort characters.Port
	}
	type args struct {
		ctx context.Context
		in1 api.CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.CreateResponse
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Campaigns{
				repository:     tt.fields.repository,
				charactersPort: tt.fields.charactersPort,
			}
			got, err := c.Create(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Campaigns.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Campaigns.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
