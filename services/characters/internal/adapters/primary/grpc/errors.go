package grpc

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AranGarcia/droop/characters/internal/core/api"

	cerr "github.com/AranGarcia/droop/shared/common-errors"
)

func handleAPIError(err error) error {
	return status.Error(apiErrToCode(err), err.Error())
}

func apiErrToCode(err error) codes.Code {
	invalidRequestError := cerr.InvalidInputError{}
	if errors.As(err, &invalidRequestError) {
		return codes.InvalidArgument
	}

	switch err {
	case api.ErrNotFound, api.ErrInvalidID:
		return codes.NotFound
	}
	return codes.Internal
}
