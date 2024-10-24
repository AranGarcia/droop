package mocks

import (
	"context"
	"fmt"

	"github.com/AranGarcia/initiatives/internal/core/entities"

	"github.com/AranGarcia/initiatives/internal/ports/core"
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
	table, ok := t.Tables[id]
	if !ok {
		return nil, core.ErrNotFound
	}
	return table, nil
}

func (t *TableRepository) Delete(_ context.Context, id string) error {
	delete(t.Tables, id)
	return nil
}
