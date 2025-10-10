package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/dnd/internal/ports/core"
	"github.com/AranGarcia/droop/dnd/internal/ports/core/mock"
	"github.com/AranGarcia/droop/dnd/internal/ports/core/rules"

	dndpb "github.com/AranGarcia/droop/protoapis/dnd/v1"
)

func TestServer_RollInitiative(t *testing.T) {
	type fields struct {
		RollInitiativeFunc func() (*rules.RollInitiativeResponse, error)
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
			fields: fields{RollInitiativeFunc: func() (*rules.RollInitiativeResponse, error) {
				return nil, core.ErrInvalidInput
			}},
			wantErr: true,
		},
		{
			name:    "internal error",
			request: &dndpb.RollInitiativeRequest{Id: "character-id"},
			fields: fields{RollInitiativeFunc: func() (*rules.RollInitiativeResponse, error) {
				return nil, errors.New("internal error")
			}},
			wantErr: true,
		},
		{
			name:    "successful initiative roll",
			request: &dndpb.RollInitiativeRequest{Id: "character-id"},
			fields: fields{RollInitiativeFunc: func() (*rules.RollInitiativeResponse, error) {
				return &rules.RollInitiativeResponse{Result: 14}, nil
			}},
			want: &dndpb.RollInitiativeResponse{Result: 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				DNDCore: mock.Rules{RollInitiativeFunc: tt.fields.RollInitiativeFunc},
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
