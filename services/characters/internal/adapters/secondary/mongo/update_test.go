package mongo

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

func levelPtr(l entities.Level) *entities.Level { return &l }

func Test_characterFieldsToBSONMap(t *testing.T) {
	tests := []struct {
		name   string
		fields repositories.CharacterFields
		want   bson.M
	}{
		{
			name:   "class",
			fields: repositories.CharacterFields{Class: repositories.StringPtr("character_class")},
			want:   bson.M{"class": "character_class"},
		},
		{
			name:   "level",
			fields: repositories.CharacterFields{Level: levelPtr(entities.Level(1))},
			want:   bson.M{"level": entities.Level(1)},
		},
		{
			name:   "name",
			fields: repositories.CharacterFields{Name: repositories.StringPtr("character_name")},
			want:   bson.M{"name": "character_name"},
		},
		{
			name:   "max_health",
			fields: repositories.CharacterFields{MaxHealth: repositories.IntPtr(1)},
			want:   bson.M{"max_health": 1},
		},
		{
			name:   "current_health",
			fields: repositories.CharacterFields{CurrentHealth: repositories.IntPtr(1)},
			want:   bson.M{"current_health": 1},
		},
		{
			name:   "temp_health",
			fields: repositories.CharacterFields{TempHealth: repositories.IntPtr(1)},
			want:   bson.M{"temp_health": 1},
		},
		{
			name:   "armor_class",
			fields: repositories.CharacterFields{ArmorClass: repositories.IntPtr(1)},
			want:   bson.M{"armor_class": 1},
		},
		{
			name:   "strength",
			fields: repositories.CharacterFields{Strength: repositories.IntPtr(1)},
			want:   bson.M{"strength": 1},
		},
		{
			name:   "dexterity",
			fields: repositories.CharacterFields{Dexterity: repositories.IntPtr(1)},
			want:   bson.M{"dexterity": 1},
		},
		{
			name:   "constitution",
			fields: repositories.CharacterFields{Constitution: repositories.IntPtr(1)},
			want:   bson.M{"constitution": 1},
		},
		{
			name:   "intelligence",
			fields: repositories.CharacterFields{Intelligence: repositories.IntPtr(1)},
			want:   bson.M{"intelligence": 1},
		},
		{
			name:   "wisdom",
			fields: repositories.CharacterFields{Wisdom: repositories.IntPtr(1)},
			want:   bson.M{"wisdom": 1},
		},
		{
			name:   "charisma",
			fields: repositories.CharacterFields{Charisma: repositories.IntPtr(1)},
			want:   bson.M{"charisma": 1},
		},
		{
			name:   "proficiencies",
			fields: repositories.CharacterFields{Proficiencies: []string{"skilla", "skillb"}},
			want:   bson.M{"proficiencies": []string{"skilla", "skillb"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterFieldsToBSONMap(tt.fields)

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("characterFieldsToBSONMap() mismatch (-want,+got):\n%s", diff)
			}
		})
	}
}
