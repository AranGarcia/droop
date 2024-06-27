package entities

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCharacter_Copy(t *testing.T) {
	now := time.Now()
	want := Character{
		Base: Base{
			ID:        "",
			CreatedAt: now,
			UpdatedAt: now,
			DeletedAt: &now,
		},
		Level:        0,
		Name:         "",
		HealthPoints: 0,
		ArmorClass:   0,
		Strength:     0,
		Dexterity:    0,
		Constitution: 0,
		Intelligence: 0,
		Wisdom:       0,
		Charisma:     0,
	}

	got := want.Copy()
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("Character.Copy() mismatch (-want,+got);\n%s", diff)
	}
}

func Test_copyTimePtr(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		t    *time.Time
		want *time.Time
	}{
		{
			name: "nil copy",
		},
		{
			name: "now",
			t:    &now,
			want: &now,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := copyTimePtr(test.t)
			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("copyTimePtr() = %v, want %v", got, test.want)
			}

			if test.want != nil && test.want != got {
				t.Fatalf("expected different pointers; got = %p, want = %p", got, test.want)
			}
		})
	}
}
