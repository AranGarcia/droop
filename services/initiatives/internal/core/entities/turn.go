package entities

// Turn based on an initiative roll
type Turn struct {
	// CharacterID is the ID of the character that the turn belongs to.
	CharacterID string
	// CharacterName is the name of the character that the turn belongs to.
	CharacterName string
	// Result of the initiative roll.
	Result int
}
