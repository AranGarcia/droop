package entities

import "github.con/AranGarcia/droop/dnd/ports/api"

// AbilityScore for a Character.
type AbilityScore struct {
	// Value of the score. It is expected to be between 1 and 30.
	Value uint
}

func NewAbilityScore(v int) (*AbilityScore, error) {
	if v < 1 || v > 30 {
		return nil, api.ErrInternalError
	}

	return &AbilityScore{Value: uint(v)}, nil
}

func (a AbilityScore) CalculateModifier() int {
	return (int(a.Value) / 2) - 5
}
