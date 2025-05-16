package models

import (
	"gorm.io/gorm"
	"time"
)

type ChatLog struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	UserID    int            `gorm:"not null" json:"user_id"`
	Keyword   string         `gorm:"type:text;not null" json:"keyword"`
	Prompt    string         `gorm:"type:text;not null" json:"prompt"`
	Response  string         `gorm:"type:text;not null" json:"response"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
