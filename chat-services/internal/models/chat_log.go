// internal/models/chat_log.go
package models

import "gorm.io/gorm"

type ChatLog struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Prompt   string `json:"prompt"`
	Response string `json:"response"`
}
