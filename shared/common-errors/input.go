package commonerrors

import "fmt"

// InvalidInputError wraps any of type of bad or invalid inputs.
type InvalidInputError struct {
	// Field that has been input.
	Field string
	// Reason of being invalid.
	Reason string
}

// Error implements the error interface.
func (i InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input for field '%s': %s", i.Field, i.Reason)
}

func (i InvalidInputError) Is(err error) bool {
	_, ok := err.(InvalidInputError)
	return ok
}
