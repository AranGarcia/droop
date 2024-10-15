package core

import "context"

type API interface {
	Initialize(context.Context) (any, error)
	AddTurn(context.Context) (any, error)
	Clear(context.Context) (any, error)
}
