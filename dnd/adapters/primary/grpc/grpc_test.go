package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"

	dndpb "github.com/AranGarcia/droop/proto/gen/dnd"
	"github.con/AranGarcia/droop/dnd/ports/core"
	"github.con/AranGarcia/droop/dnd/ports/core/mock"
)

func TestServer_RollInitiative(t *testing.T) {
	type fields struct {
		RollInitiativeFunc func() (*core.RollInitiativeResponse, error)
	}
	tests := []struct {
		name    string
		fields  fields
		request *dndpb.RollInitiativeRequest
		want    *dndpb.RollInitiativeResponse
		wantErr bool
	}{
		{
			name:    "invalid request",
			request: &dndpb.RollInitiativeRequest{},
			wantErr: true,
		},
		{
			name:    "internal error",
			request: &dndpb.RollInitiativeRequest{Id: "character-id"},
			fields: fields{RollInitiativeFunc: func() (*core.RollInitiativeResponse, error) {
				return nil, errors.New("internal error")
			}},
			wantErr: true,
		},
		{
			name:    "successful initiative roll",
			request: &dndpb.RollInitiativeRequest{Id: "character-id"},
			fields: fields{RollInitiativeFunc: func() (*core.RollInitiativeResponse, error) {
				return &core.RollInitiativeResponse{Result: 14}, nil
			}},
			want: &dndpb.RollInitiativeResponse{Result: 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				DNDCore: mock.DND{RollInitiativeFunc: tt.fields.RollInitiativeFunc},
			}

			ctx := context.Background()
			got, err := s.RollInitiative(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.RollInitiative() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.RollInitiative() = %v, want %v", got, tt.want)
			}
		})
	}
}
