package entities

type InitiativeTable struct {
	ID         string
	CampaignID string
	Turns      []Turn
}

type Turn struct {
	CharacterID   string
	CharcaterName string
	Result        int
}
