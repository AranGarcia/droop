package services

import (
	"context"
	"reflect"
	"testing"

	"github.con/AranGarcia/droop/dnd/internal/core/entities"
	"github.con/AranGarcia/droop/dnd/internal/ports/core"
	"github.con/AranGarcia/droop/dnd/internal/ports/external/characters/mock"
)

// Mock implementation of a D20 dice roll.
type mockD20 struct {
	expectedResult int
}

func (m *mockD20) Roll() int { return m.expectedResult }

func TestDND_RollInitiative(t *testing.T) {
	dex, _ := entities.NewAbilityScore(14)
	character := entities.Character{
		ID:        "character-id",
		Dexterity: *dex,
	}
	characters := mock.Characters{
		Data: map[string]entities.Character{character.ID: character},
	}

	type fields struct {
		d20 entities.Die
	}
	tests := []struct {
		name    string
		fields  fields
		request core.RollInitiativeRequest
		want    *core.RollInitiativeResponse
		wantErr bool
	}{
		{
			name:    "failed to retrieve character",
			request: core.RollInitiativeRequest{ID: "character-that-doesnt-exist"},
			wantErr: true,
		},
		{
			name: "failed to roll",
			fields: fields{
				d20: &mockD20{-1},
			},
			request: core.RollInitiativeRequest{ID: character.ID},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				d20: &mockD20{13},
			},
			request: core.RollInitiativeRequest{ID: character.ID},
			want: &core.RollInitiativeResponse{
				Result: 15,
			},
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		d := DND{
			characters: characters,
			d20:        tt.fields.d20,
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := d.RollInitiative(ctx, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("DND.RollInitiative() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DND.RollInitiative() = %v, want %v", got, tt.want)
			}
		})
	}
}
