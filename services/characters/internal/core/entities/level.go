package entities

// Level of a character.
type Level int

// CalculateProficiencyBonus returns the proficiency bonus according to the level.
func (l Level) CalculateProficiencyBonus() int {
	return 2 + (int(l)-1)/4
}
