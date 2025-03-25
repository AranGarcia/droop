package api

import "context"

type Campaigns interface {
	Create(context.Context, CreateRequest) (*CreateResponse, error)
	Retrieve(context.Context, RetrieveRequest) (*RetrieveResponse, error)
	Update(context.Context, UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, DeleteRequest) (*DeleteResponse, error)
	List(context.Context, ListRequest) (*ListResponse, error)
}

type CreateRequest struct{}
type CreateResponse struct{}
type RetrieveRequest struct{}
type RetrieveResponse struct{}
type UpdateRequest struct{}
type UpdateResponse struct{}
type DeleteRequest struct{}
type DeleteResponse struct{}
type ListRequest struct{}
type ListResponse struct{}
