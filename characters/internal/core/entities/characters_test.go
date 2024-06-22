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
	want := time.Now()
	got := copyTimePtr(&want)
	if want.Compare(*got) != 0 {
		t.Fatalf("time.Compare() mismatch; want = '%s', got '%s'", want, got)
	}
}
