package entities

import (
	"crypto/rand"
	"math/big"

	"github.com/AranGarcia/droop/dnd/internal/ports/core"
)

var (
	maxD20 = big.NewInt(20)
)

// Die is a random generator for dice rolls.
type Die interface {
	// Roll the dice and returns a number within the range [1,N].
	Roll() (int, error)
}

// D20 is a 20 faced die.
type D20 struct{}

// Roll the D20. If an error occurs, then it also returns -1.
func (d D20) Roll() (int, error) {
	roll, err := rand.Int(rand.Reader, maxD20)
	if err != nil {
		return -1, core.InternalError{
			Message: "failed to generate random roll",
			Err:     err,
		}
	}
	result := int(roll.Int64()) + 1
	return result, nil
}
