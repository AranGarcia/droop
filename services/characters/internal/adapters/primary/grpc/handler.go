package grpc

import (
	"context"

	characterspb "github.com/AranGarcia/droop/protoapis/characters/v1"
)

func (s Server) Create(ctx context.Context, request *characterspb.CreateRequest) (*characterspb.CreateResponse, error) {
	apiRequest := CreateRequestToAPI(request)
	apiResponse, err := s.characterService.Create(ctx, apiRequest)
	if err != nil {
		return nil, handleAPIError(err)
	}
	response := CreateResponseToProto(apiResponse)
	return response, nil
}

// Retrieve a Character.
func (s Server) Retrieve(ctx context.Context, request *characterspb.RetrieveRequest) (*characterspb.RetrieveResponse, error) {
	apiRequest := RetrieveRequestToAPI(request)
	apiResponse, err := s.characterService.Retrieve(ctx, apiRequest)
	if err != nil {
		return nil, handleAPIError(err)
	}
	response := RetrieveResponseToProto(apiResponse)
	return response, nil
}

// List returns all Characters.
func (s Server) List(ctx context.Context, request *characterspb.ListRequest) (*characterspb.ListResponse, error) {
	apiRequest := ListRequestToAPI(request)
	apiResponse, err := s.characterService.List(ctx, apiRequest)
	if err != nil {
		return nil, handleAPIError(err)
	}
	response := ListResponseToProto(apiResponse)
	return response, nil
}
