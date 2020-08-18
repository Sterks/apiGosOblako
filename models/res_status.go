package models

// ResStatus Response статуса заявок
type ResStatus struct {
	Result            string `json:"result"`            // error/success
	SourceDescription string `json:"sourcedescription"` // Oписание статуса
	InternalID        string `json:"internal_id"`       // Номер заявки
	Status            string `json:"status"`            // Статус заявки
	Message           string `json:"message"`           // Описание ошибки
}
