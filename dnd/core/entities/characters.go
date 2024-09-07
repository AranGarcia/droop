package entities

// Character of a player.
type Character struct {
	// Name of the character
	Name string `json:"name"`
	// Dexterity measures the agility.
	Dexterity int `json:"dexterity"`
}
