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

func TestCharacter_CalculateArmorClass(t *testing.T) {
	tests := []struct {
		name      string
		character Character
		want      int
	}{
		{
			name: "no armor and no unarmored defense",
			character: Character{
				Dexterity: AbilityScore(12),
			},
			want: 11,
		},
		{
			name: "barbarian with no armor",
			character: Character{
				Class:        BarbarianClass,
				Dexterity:    AbilityScore(13), // +1
				Constitution: AbilityScore(15), // +2
			},
			want: 13,
		},
		{
			name: "barbarian with shield but no armor",
			character: Character{
				Class:        BarbarianClass,
				Dexterity:    AbilityScore(13), // +1
				Constitution: AbilityScore(15), // +2
				Shield:       true,
			},
			want: 15,
		},
		{
			name: "monk with no armor",
			character: Character{
				Class:     MonkClass,
				Dexterity: AbilityScore(13), // +1
				Wisdom:    AbilityScore(15), // +2
			},
			want: 13,
		},
		{
			name: "monk with shield but no armor",
			character: Character{
				Class:     MonkClass,
				Dexterity: AbilityScore(13), // +1
				Wisdom:    AbilityScore(15), // +2
				Shield:    true,
			},
			want: 11,
		},
		{
			name: "armor with no shield",
			character: Character{
				Dexterity: AbilityScore(15), // +2
				Armor:     NewArmor(StudedLeatherArmor),
			},
			want: 14,
		},
		{
			name: "armor with shield",
			character: Character{
				Dexterity: AbilityScore(15),         // +2
				Armor:     NewArmor(HalfPlateArmor), // 15 AC bonus
				Shield:    true,
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.character.CalculateArmorClass(); got != tt.want {
				t.Errorf("Character.CalculateArmorClass() = %v, want %v", got, tt.want)
			}
		})
	}
}
