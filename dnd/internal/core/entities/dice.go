package entities

import (
	"crypto/rand"
	"math/big"
)

var (
	maxD20 = big.NewInt(20)
)

// Die is a random generator for dice rolls.
type Die interface {
	// Roll the dice and returns a number within the range [1,N].
	Roll() int
}

// D20 is a 20 faced die.
type D20 struct{}

// Roll the D20
func (d D20) Roll() int {
	roll, err := rand.Int(rand.Reader, maxD20)
	if err != nil {
		return -1
	}
	return int(roll.Int64()) + 1
}
