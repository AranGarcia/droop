package entities

import (
	"errors"
	"math"
	"math/rand/v2"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/dnd/internal/ports/core"
)

// randRange generates a random int within the interval [min, max].
func randRange(min, max int) int {
	interval := max - min
	return rand.IntN(interval) + min
}

func TestNewAbilityScore(t *testing.T) {
	value := randRange(1, 30)
	tests := []struct {
		name    string
		v       int
		want    *AbilityScore
		wantErr error
	}{
		{
			name:    "less than 1",
			v:       randRange(math.MinInt+1, 0),
			wantErr: core.ErrInvalidInput,
		},
		{
			name:    "more than 30",
			v:       randRange(30, math.MaxInt),
			wantErr: core.ErrInvalidInput,
		},
		{
			name: "success",
			v:    value,
			want: &AbilityScore{Value: uint(value)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAbilityScore(tt.v)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("NewAbilityScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAbilityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbilityScore_CalculateModifier(t *testing.T) {
	in := [30]int{}
	for i := range in {
		in[i] = i + 1
	}
	want := [30]int{-5, -4, -4, -3, -3, -2, -2, -1, -1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10}
	for i, v := range in {
		a := AbilityScore{Value: uint(v)}
		got := a.CalculateModifier()

		if got != want[i] {
			t.Errorf("AbilityScore.CalculateModifier(%d) = %d, want %d", v, got, want[i])
		}
	}
}
