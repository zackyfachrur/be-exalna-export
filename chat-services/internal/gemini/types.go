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

type ProductProps struct {
	CompanyName     *string  `json:"companyName,omitempty"`
	UserName        *string  `json:"userName,omitempty"`
	ThumbnailURL    *string  `json:"thumbnailURL,omitempty"`
	ProductName     *string  `json:"productName,omitempty"`
	IsImport        *bool    `json:"isImport,omitempty"`
	IsExport        *bool    `json:"isExport,omitempty"`
	ProductCategory *string  `json:"productCategory,omitempty"`
	Quantity        *int     `json:"quantity,omitempty"`
	Cost            *float64 `json:"cost,omitempty"`
}
