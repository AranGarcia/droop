package characters

import (
	"context"
	"errors"
	"reflect"
	"testing"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
)

func TestClient_Retrieve(t *testing.T) {
	type fields struct {
		retrieveFunc func() (*characterspb.RetrieveResponse, error)
	}
	tests := []struct {
		name    string
		fields  fields
		want    *entities.Character
		wantErr bool
	}{
		{
			name: "client error",
			fields: fields{
				retrieveFunc: func() (*characterspb.RetrieveResponse, error) {
					return nil, errors.New("mock retrieve error")
				},
			},
			wantErr: true,
		},
		{
			name: "successful retrieve",
			fields: fields{
				retrieveFunc: func() (*characterspb.RetrieveResponse, error) {
					response := &characterspb.RetrieveResponse{
						Character: &characterspb.Character{
							Id:        "character-id",
							Dexterity: 10,
						},
					}
					return response, nil
				},
			},
			want: &entities.Character{
				ID:        "character-id",
				Dexterity: entities.AbilityScore{Value: 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				grpcClient: mockClient{retrieveFunc: tt.fields.retrieveFunc},
			}

			ctx := context.Background()
			got, err := c.Retrieve(ctx, "mock-id")
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}
