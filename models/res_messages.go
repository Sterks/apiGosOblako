package models

//ResMessages ...
type ResMessages struct {
	Owner      string `json:"owner"`
	InternalID string `json:"internal_id"`
	Source     string `json:"source"`
	ID         string `json:"id"`
	Status     string `json:"status"`
	StatusCode string `json:"statusCode"`
}
