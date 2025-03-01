package mocks

import (
	"context"
	"errors"
	"testing"

	"github.com/AranGarcia/droop/characters/internal/core/entities"
	"github.com/AranGarcia/droop/characters/internal/ports/repositories"
	"github.com/google/go-cmp/cmp"
)

// TestImplementation is expected to throw a compile error rather than a runtime assertion to
// validate that the mock is an implementation of the expected repository.
func TestImplementation(t *testing.T) {
	var _ repositories.Characters = &Characters{}
}

func TestCharacters_Create(t *testing.T) {
	character := entities.Character{
		Base: entities.Base{ID: "pre-existing-character"},
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
			name: "successful create",
			in:   entities.Character{Base: entities.Base{ID: "completely-new-character"}},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Characters{inMemory: inMemory}
			_, err := c.Create(ctx, tt.in)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Characters.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCharacters_Retrieve(t *testing.T) {
	character := entities.Character{
		Base: entities.Base{ID: "pre-existing-character"},
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
	characterID := "character-id"
	type fields struct {
		inMemory map[string]entities.Character
	}
	tests := []struct {
		name            string
		id              string
		fields          fields
		characterFields repositories.CharacterFields
		want            *entities.Character
		wantErr         error
	}{
		{
			name:    "not found",
			id:      "not-found-character-id",
			wantErr: repositories.ErrNotFound,
		},
		{
			name:            "level",
			id:              characterID,
			characterFields: repositories.CharacterFields{Level: repositories.IntPtr(10)},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, Level: 10},
		},
		{
			name:            "name",
			id:              characterID,
			characterFields: repositories.CharacterFields{Name: repositories.StringPtr("new-character-name")},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, Name: "new-character-name"},
		},
		{
			name:            "armorclass",
			id:              characterID,
			characterFields: repositories.CharacterFields{ArmorClass: repositories.IntPtr(15)},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, ArmorClass: 15},
		},
		{
			name:            "strength",
			id:              characterID,
			characterFields: repositories.CharacterFields{Strength: repositories.IntPtr(15)},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, Strength: 15},
		},
		{
			name:            "dexterity",
			id:              characterID,
			characterFields: repositories.CharacterFields{Dexterity: repositories.IntPtr(15)},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, Dexterity: 15},
		},
		{
			name:            "constitution",
			id:              characterID,
			characterFields: repositories.CharacterFields{Constitution: repositories.IntPtr(15)},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, Constitution: 15},
		},
		{
			name:            "intelligence",
			id:              characterID,
			characterFields: repositories.CharacterFields{Intelligence: repositories.IntPtr(15)},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, Intelligence: 15},
		},
		{
			name:            "wisdom",
			id:              characterID,
			characterFields: repositories.CharacterFields{Wisdom: repositories.IntPtr(15)},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, Wisdom: 15},
		},
		{
			name:            "charisma",
			id:              characterID,
			characterFields: repositories.CharacterFields{Charisma: repositories.IntPtr(15)},
			fields:          fields{map[string]entities.Character{characterID: {Base: entities.Base{ID: characterID}}}},
			want:            &entities.Character{Base: entities.Base{ID: characterID}, Charisma: 15},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Characters{inMemory: tt.fields.inMemory}
			got, err := c.Update(ctx, tt.id, tt.characterFields)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Characters.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Characters.Update() mismatch (-want,+got):\n%s", diff)
			}
		})
	}
}

func TestCharacters_Delete(t *testing.T) {
	type fields struct {
		inMemory map[string]entities.Character
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]entities.Character
		wantErr bool
	}{
		{
			name: "no delete",
			args: args{
				id: "non-existing-character-id",
			},
			fields: fields{
				inMemory: map[string]entities.Character{
					"non-affected-id": {Base: entities.Base{ID: "non-affected-id"}},
				},
			},
			want: map[string]entities.Character{
				"non-affected-id": {Base: entities.Base{ID: "non-affected-id"}},
			},
		},
		{
			name: "successful delete",
			args: args{
				id: "expected-to-delete",
			},
			fields: fields{
				inMemory: map[string]entities.Character{
					"expected-to-delete": {Base: entities.Base{ID: "expected-to-delete"}},
				},
			},
			want: map[string]entities.Character{},
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Characters{
				inMemory: tt.fields.inMemory,
			}
			_ = c.Delete(ctx, tt.args.id) // mock doesn't throw errors
			if diff := cmp.Diff(tt.want, c.inMemory); diff != "" {
				t.Errorf("mismatch(-want,+got):\n%s", diff)
			}
		})
	}
}
