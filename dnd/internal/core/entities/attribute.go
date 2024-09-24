package entities

import (
	"errors"
)

// AbilityScore for a Character.
type AbilityScore struct {
	// Value of the score. It is expected to be between 1 and 30.
	Value uint
}

func NewAbilityScore(v int) (*AbilityScore, error) {
	if v < 1 || v > 30 {
		return nil, errors.New("ability score must be in range [1, 30]")
	}

	return &AbilityScore{Value: uint(v)}, nil
}

func (a AbilityScore) CalculateModifier() int {
	return (int(a.Value) / 2) - 5
}
