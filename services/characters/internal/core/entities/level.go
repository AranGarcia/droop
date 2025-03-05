package entities

// Level of a character.
type Level int

// NewLevel creates a new level from 32 bit signed integer.
func NewLevelFrom32(l int32) Level {
	return Level(l)
}

// CalculateProficiencyBonus returns the proficiency bonus according to the level.
func (l Level) CalculateProficiencyBonus() int {
	return 2 + (int(l)-1)/4
}
