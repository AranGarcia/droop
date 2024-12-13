// Package rules is the API for the D&D interactions and rules for playing.
package rules

import "context"

type API interface {
	RollInitiative(context.Context, RollInitiativeRequest) (*RollInitiativeResponse, error)
}

// RollInitiativeRequest for API.RollInitiative.
type RollInitiativeRequest struct {
	// ID of the character that will roll initiative.
	ID string
}

// RollInitiativeResponse of API.RollInitiative.
type RollInitiativeResponse struct {
	// Result of the initiative roll.
	Result int
}
