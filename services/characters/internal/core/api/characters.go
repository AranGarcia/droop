package api

import (
	"context"
)

type Characters interface {
	Create(context.Context, CreateCharacterRequest) (*CreateCharacterResponse, error)
	Retrieve(context.Context, RetrieveCharacterRequest) (*RetrieveCharacterResponse, error)
	Update(context.Context, UpdateCharacterRequest) (*UpdateCharacterResponse, error)
	Delete(context.Context, DeleteCharacterRequest) (*DeleteCharacterResponse, error)
	List(context.Context, ListCharactersRequest) (*ListCharactersResponse, error)
	GenerateSheet(context.Context, GenerateSheetRequest) (*GenerateSheetResponse, error)
}
