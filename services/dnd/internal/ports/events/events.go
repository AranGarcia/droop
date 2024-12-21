// events is an outgoing port that produces domain events.
package events

import "context"

// Producer of domain events.
type Port interface {
	// RollInitiativeSuccess produces the event of a successful execution of RollInitiative.
	RollInitiativeSuccess(context.Context, RollInitiativeSuccessMessage) error
}

// RollInitiativeSuccessMessage contains information on the result of an initiative roll.
type RollInitiativeSuccessMessage struct {
	// CharacterID that rolled for initiative.
	CharacterID string
	// Result of the roll.
	Result int
}
