package entities

import (
	"math/rand/v2"
	"testing"
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
	proficiencies := []string{
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
		Abilities: map[string]Skill{
			proficiencies[rand.IntN(len(proficiencies))]: {
				Proficient: rand.IntN(2) == 1,
				Bonus:      rand.IntN(4),
			},
			proficiencies[rand.IntN(len(proficiencies))]: {
				Proficient: rand.IntN(2) == 1,
				Bonus:      rand.IntN(4),
			},
			proficiencies[rand.IntN(len(proficiencies))]: {
				Proficient: rand.IntN(2) == 1,
				Bonus:      rand.IntN(4),
			},
		},
	}
	if err := character.Validate(); err != nil {
		t.Fatalf("validation error; %v", err)
	}
}
