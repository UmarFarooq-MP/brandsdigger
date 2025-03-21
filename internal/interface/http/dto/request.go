package dto

type NamesRequest struct {
	Message string `json:"message"`
}

type NamesRequestWithContext struct {
	NewMessage string   `json:"new_message"`
	History    []string `json:"history"`
}
