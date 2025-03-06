package entities

import (
	"testing"
)

func TestAbilityScore_Modifier(t *testing.T) {
	in := [30]int{}
	for i := range in {
		in[i] = i + 1
	}
	want := [30]int{-5, -4, -4, -3, -3, -2, -2, -1, -1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10}

	for i, v := range in {
		a := AbilityScore(v)

		if got := a.Modifier(); got != want[i] {
			t.Errorf("AbilityScore.Modifier(%d) = %d, want %d", v, got, want[i])
		}
	}
}
