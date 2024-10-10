package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
	"github.com/AranGarcia/droop/dnd/internal/ports/core"
	eventsmock "github.com/AranGarcia/droop/dnd/internal/ports/events/mock"
	charactersmock "github.com/AranGarcia/droop/dnd/internal/ports/external/characters/mock"
)

// Mock implementation of a D20 dice roll.
type mockD20 struct {
	rollFunc func() (int, error)
}

func (m mockD20) Roll() (int, error) { return m.rollFunc() }

func TestDND_RollInitiative(t *testing.T) {
	dex, _ := entities.NewAbilityScore(14)
	character := entities.Character{
		ID:        "character-id",
		Dexterity: *dex,
	}
	characters := charactersmock.Characters{
		Data: map[string]entities.Character{character.ID: character},
	}

	type fields struct {
		d20      entities.Die
		producer eventsmock.Producer
	}
	tests := []struct {
		name               string
		fields             fields
		request            core.RollInitiativeRequest
		want               *core.RollInitiativeResponse
		wantProducedEvents int
		wantErr            bool
	}{
		{
			name:    "empty ID",
			request: core.RollInitiativeRequest{},
			wantErr: true,
		},
		{
			name:    "failed to retrieve character",
			request: core.RollInitiativeRequest{ID: "character-that-doesnt-exist"},
			wantErr: true,
		},
		{
			name: "failed to roll",
			fields: fields{
				d20: &mockD20{rollFunc: func() (int, error) { return -1, errors.New("random gen error") }},
			},
			request: core.RollInitiativeRequest{ID: character.ID},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				d20:      &mockD20{rollFunc: func() (int, error) { return 13, nil }},
				producer: eventsmock.Producer{},
			},
			request: core.RollInitiativeRequest{ID: character.ID},
			want: &core.RollInitiativeResponse{
				Result: 15,
			},
			wantProducedEvents: 1,
		},
	}
	ctx := context.Background()
	for _, tt := range tests {
		d := DND{
			characters: characters,
			d20:        tt.fields.d20,
			events:     &tt.fields.producer,
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

			gotProducedEvents := tt.fields.producer.AmountOfMessagesProduced()
			if gotProducedEvents != tt.wantProducedEvents {
				t.Errorf("DND.RollInitiative() produced events = %d, want %d", gotProducedEvents, tt.wantProducedEvents)
			}
		})
	}
}
