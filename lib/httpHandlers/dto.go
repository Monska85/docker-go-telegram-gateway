package httpHandlers

type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type SendMessageRequest struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
