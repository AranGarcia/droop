package core

// TODO: figure out if this needs to move to its own package.

import "context"

type DND interface {
	RollInitiative(context.Context, RollInitiativeRequest) (*RollInitiativeResponse, error)
}

// RollInitiativeRequest for DND.RollInitiative.
type RollInitiativeRequest struct {
	// ID of the character that will roll initiative.
	ID string
}

// RollInitiativeResponse of DND.RollInitiative.
type RollInitiativeResponse struct {
	// Result of the initiative roll.
	Result int
}
