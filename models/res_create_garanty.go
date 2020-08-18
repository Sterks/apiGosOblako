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
	// Messages   Mess   `json:"message"`
}

// SubMessageData ...
type SubMessageData struct {
	FormatValidation ReqCreateGaranty `json:"formatValidation"`
}

// FormatValidation ...
type FormatValidation struct {
	Message Mess `json:"message"`
}

// Mess ...
type Mess struct {
	Message string `json:"message"`
}
