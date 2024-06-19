package mocks

import (
	"context"
	"errors"
	"testing"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
)

func TestCharacters_Create(t *testing.T) {
	character := entities.Character{
		ID: "pre-existing-character",
	}
	inMemory := map[string]entities.Character{
		character.ID: character,
	}
	tests := []struct {
		name    string
		in      entities.Character
		wantErr error
	}{
		{
			name:    "duplicate",
			in:      entities.Character{ID: character.ID},
			wantErr: repositories.ErrDuplicateEntity,
		},
		{
			name: "successful create",
			in:   entities.Character{ID: "completely-new-character"},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Characters{inMemory: inMemory}
			err := c.Create(ctx, tt.in)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Characters.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
