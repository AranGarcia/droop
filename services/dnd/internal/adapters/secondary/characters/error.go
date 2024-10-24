package characters

import "errors"

var (
	ErrCharacterServiceError = errors.New("character service error")
	ErrInvalidAttributeType  = errors.New("invalid attribute type")
)
