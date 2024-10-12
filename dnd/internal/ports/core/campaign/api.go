// Package campaign contains a port to the core service for Campaigns.
package campaign

import (
	"context"

	"github.com/AranGarcia/droop/dnd/internal/core/entities"
)

// API definition for Campaigns.
type API interface {
	// Create a Campaign.
	Create(context.Context, CreateRequest) (*CreateResponse, error)
	// Retrieve a Campaign.§
	Retrieve(context.Context, RetrieveRequest) (*RetrieveResponse, error)
}

// CreateRequest is a request for the method API.Create
type CreateRequest struct {
	// Name of the campaign to be created.
	Name string
}

// CreateResponse is a response for the method API.Create
type CreateResponse struct {
	Campaign entities.Campaign
}

// RetrieveRequest is a request for the method API.Retrieve
type RetrieveRequest struct {
	// ID of the campaign to be retrieved.
	ID string
}

// RetrieveResponse is a response for the method API.Retrieve
type RetrieveResponse struct {
	Campaign entities.Campaign
}
