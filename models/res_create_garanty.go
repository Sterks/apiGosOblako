package models

// ResCreateGaranty ...
type ResCreateGaranty struct {
	Messa      FounderFLData
	ID         int    `json:"id"` // Идентификатор записи в таблице ai_deals
	ExternalID string `json:"external_id"`
	Source     string `json:"source"`
	Status     string `json:"status"`
	Owner      string `json:"owner"`
	InternalID string `json:"internal_id"`
}

// SubMessageData ...
type SubMessageData struct {
	FormatValidation ReqCreateGaranty
}

// Mess ...
type Mess struct {
	Message string `json:"message"`
}
