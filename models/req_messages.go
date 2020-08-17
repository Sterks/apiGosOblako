package models

import "time"

// ReqMessages ...
type ReqMessages struct {
	InternalID string `json:"internal_id"`
	ResultM    ResultM
}

// ResultM ...
type ResultM struct {
	Messages1 Messages
	Messages2 Messages
}

// Messages ...
type Messages struct {
	Date time.Time `json:"date"`
	ID   string    `json:"id"`
	Mess string    `json:"mess"`
	User string    `json:"user"`
}
