package api

import "context"

type DND interface {
	RollInitiative(context.Context, RollInitiativeRequest) (*RollInitiativeResponse, error)
}

type RollInitiativeRequest struct{}
type RollInitiativeResponse struct{}
