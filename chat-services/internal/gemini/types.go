package gemini

type Service struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

type GeminiRequest struct {
	UserID uint   `json:"user_id"`
	Prompt string `json:"prompt"`
}
