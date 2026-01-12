package dto

// Struct email request
type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	Name    string `json:"name"`
}
