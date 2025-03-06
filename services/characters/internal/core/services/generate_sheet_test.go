package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/characters/internal/core/api"
	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories/mocks"
)

func TestCharacters_GenerateSheet(t *testing.T) {
	tests := []struct {
		name       string
		repository repositories.Characters
		request    api.GenerateSheetRequest
		want       *api.GenerateSheetResponse
		wantErr    bool
	}{
		{
			name:       "not found",
			request:    api.GenerateSheetRequest{ID: "not-found"},
			repository: &mocks.Characters{},
			wantErr:    true,
		},
		{
			name: "calculated success",
			request: api.GenerateSheetRequest{
				ID: "character-id",
			},
			repository: &mocks.Characters{
				InMemory: map[string]entities.Character{
					"character-id": {
						Level:     entities.Level(6),
						Dexterity: entities.AbilityScore(15),
					},
				},
			},
			want: &api.GenerateSheetResponse{
				Characer: entities.Character{
					Level:     entities.Level(6),
					Dexterity: entities.AbilityScore(15),
				},
				Proficiency: 3,
				Initiative:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Characters{
				repository: tt.repository,
			}
			ctx := context.Background()
			got, err := c.GenerateSheet(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Characters.GenerateSheet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Characters.GenerateSheet() = %v, want %v", got, tt.want)
			}
		})
	}
}
