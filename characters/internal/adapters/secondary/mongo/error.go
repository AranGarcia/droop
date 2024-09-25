package mongo

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrorNotFound    = errors.New("not found")
	ErrInternalError = errors.New("unknown error")
	ErrInvalidID     = errors.New("invalid ObjectID")
)

func handleMongoError(err error) error {
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return ErrorNotFound
	}
	return ErrInternalError
}
