package api

import "github.com/AranGarcia/droop/characters/internal/core/entities"

type GenerateSheetRequest struct {
	// ID of the character to generate the sheet for.
	ID string
}

type GenerateSheetResponse struct {
	Characer          entities.Character
	PassivePerception int
	ArmorClass        int
	Speed             int

	ProficiencyBonus int
	InitiativeBonus  int

	// Strength-based skills

	Athletics int

	// Dexterity-based skills

	Acrobatics    int
	SleightOfHand int
	Stealth       int

	// Intelligence-based skills

	Arcana        int
	History       int
	Investigation int
	Nature        int
	Religion      int

	// Wisdom-based skills

	AnimalHandling int
	Insight        int
	Medicine       int
	Perception     int
	Survival       int

	// Charisma-based skills

	Deception    int
	Intimidation int
	Performance  int
	Persuasion   int
}
