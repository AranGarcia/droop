package grpc

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AranGarcia/droop/characters/internal/ports/api"
)

func handleAPIError(err error) error {
	return status.Error(apiErrToCode(err), err.Error())
}

func apiErrToCode(err error) codes.Code {
	invalidRequestError := api.InvalidRequestError{}
	if errors.As(err, &invalidRequestError) {
		return codes.InvalidArgument
	}

	switch err {
	case api.ErrNotFound:
		return codes.NotFound
	}
	return codes.Internal
}
