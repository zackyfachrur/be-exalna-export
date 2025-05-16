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
	UserID  int    `json:"user_id"`
	Keyword string `json:"keyword"`
	Prompt  string `json:"prompt"`
}
