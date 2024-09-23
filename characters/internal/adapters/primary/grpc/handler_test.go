package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/api/mock"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

func TestServer_Retrieve(t *testing.T) {
	character := entities.Character{Base: entities.Base{ID: "character-id"}}
	tests := []struct {
		name    string
		request *characterspb.RetrieveRequest
		want    *characterspb.RetrieveResponse
		wantErr bool
	}{
		{
			name:    "invalid request",
			request: &characterspb.RetrieveRequest{},
			wantErr: true,
		},
		{
			name:    "service error",
			request: &characterspb.RetrieveRequest{Id: "character that doesn't exist"},
			wantErr: true,
		},
		{
			name:    "successful retrieve",
			request: &characterspb.RetrieveRequest{Id: character.ID},
			want: &characterspb.RetrieveResponse{
				Character: &characterspb.Character{
					Id: character.ID,
				},
			},
		},
	}

	mockService := mock.Characters{
		Entities: map[string]entities.Character{
			"character-id": character,
		},
	}
	s := Server{characterService: &mockService}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Retrieve(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}
