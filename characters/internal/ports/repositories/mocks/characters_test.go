package mocks

import (
	"context"
	"errors"
	"testing"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
	"github.com/google/go-cmp/cmp"
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

func TestCharacters_Retrieve(t *testing.T) {
	character := entities.Character{
		ID: "pre-existing-character",
	}
	inMemory := map[string]entities.Character{
		character.ID: character,
	}
	tests := []struct {
		name    string
		in      string
		want    *entities.Character
		wantErr error
	}{
		{
			name:    "not found",
			in:      "non-existing-character-id",
			wantErr: repositories.ErrNotFound,
		},
		{
			name: "successful retrieve",
			in:   character.ID,
			want: &character,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Characters{inMemory: inMemory}
			got, err := c.Retrieve(ctx, tt.in)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Characters.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Retrieve() mismatch (-want, +got):\n%s", diff)
			}
		})

	}
}

func TestCharacters_Update(t *testing.T) {
	type fields struct {
		inMemory map[string]entities.Character
	}
	type args struct {
		in0 context.Context
		in1 entities.Character
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Characters{
				inMemory: tt.fields.inMemory,
			}
			if err := c.Update(tt.args.in0, tt.args.in1); (err != nil) != tt.wantErr {
				t.Errorf("Characters.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
