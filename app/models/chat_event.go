package models

type ChatEvent struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	PartyID   string `json:"partyId"`
	Type      string `json:"type"` // "join", "leave", or "message"
	UserID    string `json:"userId"`
	Timestamp int    `json:"timestamp"` // Unix timestamp (secs)
	Text      string `json:"text"`      // What the user said (if Type == "message")
}
