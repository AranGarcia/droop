package grpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AranGarcia/droop/initiatives/internal/ports/core"
)

func handleAPIError(err error) error {
	return status.Error(apiErrToCode(err), err.Error())
}

func apiErrToCode(err error) codes.Code {
	switch err {
	case core.ErrInvalidID:
		return codes.InvalidArgument
	case core.ErrNotFound:
		return codes.NotFound
	}
	return codes.Internal
}
