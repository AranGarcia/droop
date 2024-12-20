package characters

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound             = errors.New("character not found")
	ErrInvalidAttributeType = errors.New("invalid attribute type")
	ErrUnexpectedError      = errors.New("unexpected error")
)

func translateGRPCError(err error) error {
	if err == nil {
		return err
	}

	status := status.Convert(err)
	if status == nil {
		return ErrUnexpectedError
	}

	switch status.Code() {
	case codes.NotFound:
		return ErrNotFound
	}
	return ErrUnexpectedError
}
