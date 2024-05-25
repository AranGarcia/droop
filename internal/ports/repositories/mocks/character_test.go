package mocks

import (
	"context"
	"reflect"
	"testing"

	"github.com/AranGarcia/droop/internal/core/entities"
	"github.com/AranGarcia/droop/internal/ports"
	"github.com/AranGarcia/droop/internal/ports/repositories"
)

func TestCharacterMockRepository_Update(t *testing.T) {
	type fields struct {
		characters map[string]entities.Character
	}
	type args struct {
		id string
		in repositories.UpdateCharacterInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Character
		wantErr bool
	}{
		{
			name:    "not found",
			wantErr: true,
		},
		{
			name: "non-nil name",
			fields: fields{
				characters: map[string]entities.Character{
					"non-nil-name-character": {},
				},
			},
			args: args{
				id: "non-nil-name-character",
				in: repositories.UpdateCharacterInput{
					Name: ports.StringPtr("new-name"),
				},
			},
			want: &entities.Character{Name: "new-name"},
		},
		{
			name: "non-nil level",
			fields: fields{
				characters: map[string]entities.Character{
					"non-nil-level-character": {},
				},
			},
			args: args{
				id: "non-nil-level-character",
				in: repositories.UpdateCharacterInput{
					Level: ports.UintPtr(20),
				},
			},
			want: &entities.Character{Level: 20},
		},
		{
			name: "non-nil hit points",
			fields: fields{
				characters: map[string]entities.Character{
					"non-nil-hitpoints-character": {},
				},
			},
			args: args{
				id: "non-nil-hitpoints-character",
				in: repositories.UpdateCharacterInput{
					HitPoints: ports.IntPtr(-10),
				},
			},
			want: &entities.Character{HitPoints: -10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CharacterMockRepository{
				characters: tt.fields.characters,
			}
			got, err := c.Update(context.Background(), tt.args.id, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CharacterMockRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CharacterMockRepository.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
