package entities

// AbilityScore represents a character's different capabilities.
type AbilityScore int

// Modifier calculated for the ability score.
func (a AbilityScore) Modifier() int {
	return (int(a) / 2) - 5
}
