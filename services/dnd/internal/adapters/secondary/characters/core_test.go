package characters

import (
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"

	characterspb "github.com/AranGarcia/droop/proto/gen/characters"
)

func TestCharacterCoreFromExternal(t *testing.T) {
	tests := []struct {
		character *characterspb.Character
		name      string
		want      *entities.Character
		wantErr   bool
	}{
		{
			name:      "invalid attribute",
			wantErr:   true,
			character: &characterspb.Character{Dexterity: 0},
		},
		{
			name:      "success character core",
			character: &characterspb.Character{Dexterity: 17},
			want:      &entities.Character{Dexterity: entities.AbilityScore{Value: 17}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CharacterCoreFromExternal(tt.character)
			if (err != nil) != tt.wantErr {
				t.Errorf("CharacterCoreFromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CharacterCoreFromExternal() = %v, want %v", got, tt.want)
			}
		})
	}
}
