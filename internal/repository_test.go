package internal

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateFields_ToBsonMap(t *testing.T) {

	tests := []struct {
		name string
		u    UpdateFields
		want bson.M
	}{
		{
			name: "level",
			u:    UpdateFields{Level: IntPtr(19)},
			want: bson.M{"level": 19},
		},
		{
			name: "name",
			u:    UpdateFields{Name: StrPtr("Frido")},
			want: bson.M{"name": "Frido"},
		},
		{
			name: "health points",
			u:    UpdateFields{HealthPoints: IntPtr(77)},
			want: bson.M{"health_points": 77},
		},
		{
			name: "armor class",
			u:    UpdateFields{ArmorClass: IntPtr(15)},
			want: bson.M{"armor_class": 15},
		},
		{
			name: "strength",
			u:    UpdateFields{Strength: IntPtr(10)},
			want: bson.M{"strength": 10},
		},
		{
			name: "dexterity",
			u:    UpdateFields{Dexterity: IntPtr(10)},
			want: bson.M{"dexterity": 10},
		},
		{
			name: "constitution",
			u:    UpdateFields{Constitution: IntPtr(10)},
			want: bson.M{"constitution": 10},
		},
		{
			name: "intelligence",
			u:    UpdateFields{Intelligence: IntPtr(10)},
			want: bson.M{"intelligence": 10},
		},
		{
			name: "wisdom",
			u:    UpdateFields{Wisdom: IntPtr(10)},
			want: bson.M{"wisdom": 10},
		},
		{
			name: "charisma",
			u:    UpdateFields{Charisma: IntPtr(10)},
			want: bson.M{"charisma": 10},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.u.ToBsonMap(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("UpdateFields.ToBsonMap() = %v, want %v", got, test.want)
			}
		})
	}
}
