package entities

import (
	"math/rand/v2"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestCharacter_Validate(t *testing.T) {
	classes := []ClassName{
		BarbarianClass,
		BardClass,
		ClericClass,
		DruidClass,
		FighterClass,
		MonkClass,
		PaladinClass,
		RangerClass,
		RogueClass,
		SorcererClass,
		WarlockClass,
		WizardClass,
	}
	proficiencies := []Skill{
		AcrobaticsSkill,
		AnimalHandlingSkill,
		ArcanaSkill,
		AthleticsSkill,
		DeceptionSkill,
		HistorySkill,
		InsightSkill,
		IntimidationSkill,
		InvestigationSkill,
		MedicineSkill,
		NatureSkill,
		PerceptionSkill,
		PerformanceSkill,
		PersuasionSkill,
		ReligionSkill,
		SleightOfHandSkill,
		StealthSkill,
		SurvivalSkill,
	}
	character := Character{
		Class:        classes[rand.IntN(len(classes))],
		Level:        Level(rand.IntN(20) + 1),
		Name:         "The Hero",
		MaxHealth:    5,
		ArmorClass:   10,
		Strength:     10,
		Dexterity:    10,
		Constitution: 10,
		Intelligence: 10,
		Wisdom:       10,
		Charisma:     10,
		Proficiencies: []Skill{
			proficiencies[rand.IntN(len(proficiencies))],
			proficiencies[rand.IntN(len(proficiencies))],
			proficiencies[rand.IntN(len(proficiencies))],
		},
	}
	if err := character.Validate(); err != nil {
		t.Fatalf("validation error; %v", err)
	}
}

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

			if test.want != nil && test.want == got {
				t.Fatalf("expected different pointers; got = %p, want = %p", got, test.want)
			}
		})
	}
}
