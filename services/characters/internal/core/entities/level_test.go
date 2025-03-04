package entities

import (
	"math/rand/v2"
	"testing"
)

func TestLevel_CalculateProficiencyBonus(t *testing.T) {
	tests := []struct {
		name string
		l    Level
		want int
	}{
		{
			name: "level 1-4",
			l:    Level(rand.IntN(4) + 1),
			want: 2,
		},
		{
			name: "level 5-8",
			l:    Level(rand.IntN(4) + 5),
			want: 3,
		},
		{
			name: "level 9-12",
			l:    Level(rand.IntN(4) + 9),
			want: 4,
		},
		{
			name: "level 13-16",
			l:    Level(rand.IntN(4) + 13),
			want: 5,
		},
		{
			name: "level 17-20",
			l:    Level(rand.IntN(4) + 17),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.CalculateProficiencyBonus(); got != tt.want {
				t.Errorf("Level.CalculateProficiencyBonus() = %v, want %v", got, tt.want)
			}
		})
	}
}
