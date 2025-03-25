package mock

import (
	"context"

	"github.com/AranGarcia/droop/campaigns/internal/core/entities"
	"github.com/AranGarcia/droop/campaigns/internal/ports/campaigns"
)

type MockRepository struct{}

func (m *MockRepository) Create(_ context.Context, _ entities.Campaign) error {
	panic("not implemented")
}

func (m *MockRepository) Retrieve(_ context.Context, _ string) (*entities.Campaign, error) {
	panic("not implemented")
}

func (m *MockRepository) Update(_ context.Context, _ campaigns.UpdateFields) (*entities.Campaign, error) {
	panic("not implemented")
}

func (m *MockRepository) Delete(_ context.Context, _ string) error {
	panic("not implemented")
}
