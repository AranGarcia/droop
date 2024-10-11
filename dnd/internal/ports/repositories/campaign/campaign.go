// Package campaign is an outgoing port to access a repository for Campaigns.
package campaign

import "context"

// Repository is an outbound port to access the repository for Campaigns.
type Repository interface {
	CreateCampaign(context.Context, CreateCampaignRequest) (*CreateCampaignResponse, error)
	RetrieveCampaign(context.Context, RetrieveCampaignRequest) (*RetrieveCampaignResponse, error)
	UpdateCampaign(context.Context, UpdateCampaignRequest) (*UpdateCampaignResponse, error)
	DeleteCampaign(context.Context, DeleteCampaignRequest) (*DeleteCampaignResponse, error)
}

type CreateCampaignRequest struct{}
type CreateCampaignResponse struct{}
type RetrieveCampaignRequest struct{}
type RetrieveCampaignResponse struct{}
type UpdateCampaignRequest struct{}
type UpdateCampaignResponse struct{}
type DeleteCampaignRequest struct{}
type DeleteCampaignResponse struct{}
