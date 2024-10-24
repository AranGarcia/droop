package entities

// Table that tracks and maintains an order for initiative turns.
type Table struct {
	// ID of the table
	ID string
	// CampaignID that the table is tracking turns of.
	CampaignID string
}
