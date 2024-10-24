package mocks

import (
	"context"
	"fmt"

	"github.com/AranGarcia/droop/initiatives/internal/core/entities"

	"github.com/AranGarcia/droop/initiatives/internal/ports/core"
)

type TableRepository struct {
	Tables map[string]*entities.Table
}

func (t *TableRepository) Create(_ context.Context, table entities.Table) error {
	_, exists := t.Tables[table.CampaignID]
	if exists {
		return fmt.Errorf("table with id %q already exists", table.CampaignID)
	}
	t.Tables[table.CampaignID] = &table
	return nil
}

func (t *TableRepository) Retrieve(_ context.Context, id string) (*entities.Table, error) {
	table, exists := t.Tables[id]
	if !exists {
		return nil, core.ErrNotFound
	}
	return table, nil
}

func (t *TableRepository) Delete(_ context.Context, id string) error {
	_, exists := t.Tables[id]
	if !exists {
		return core.ErrNotFound
	}
	delete(t.Tables, id)
	return nil
}
