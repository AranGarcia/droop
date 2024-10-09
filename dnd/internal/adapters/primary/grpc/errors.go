package grpc

import (
	"github.com/AranGarcia/droop/dnd/internal/ports/core"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func adaptCoreError(err error) error {
	return status.Error(apiErrorToCode(err), err.Error())
}

func apiErrorToCode(err error) codes.Code {
	var code codes.Code
	switch err {
	case core.ErrNoIDProvided, core.ErrInvalidInput:
		code = codes.InvalidArgument
	case core.ErrNotFound:
		code = codes.NotFound
	case core.ErrRepositoryFailure:
		code = codes.Unavailable
	case core.ErrInternalError:
		code = codes.Internal
	default:
		code = codes.Unknown
	}
	return code
}
