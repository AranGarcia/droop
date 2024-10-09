package entities

import (
	"github.com/AranGarcia/droop/dnd/internal/ports/core"
)

// AbilityScore for a Character.
type AbilityScore struct {
	// Value of the score. It is expected to be between 1 and 30.
	Value uint
}

func NewAbilityScore(v int) (*AbilityScore, error) {
	if v < 1 || v > 30 {
		return nil, core.ErrInvalidInput
	}

	return &AbilityScore{Value: uint(v)}, nil
}

func (a AbilityScore) CalculateModifier() int {
	return (int(a.Value) / 2) - 5
}
