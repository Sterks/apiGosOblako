package models

// User Структура с логином и паролем
type User struct {
	ClientID string `json:"client_id"`
	Password string `json:"password"`
}
