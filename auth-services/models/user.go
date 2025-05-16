package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Username          string         `json:"username" gorm:"unique;not null;size:255"`
	Email             string         `json:"email" gorm:"unique;not null;size:255"`
	CompanyName       string         `json:"companyName" gorm:"not null;size:255"`
	CompanyCategories string         `json:"companyCategories" gorm:"not null;size:200"`
	YearsOfExperience int            `json:"yearsOfExperience" gorm:"not null"`
	Password          string         `json:"password" gorm:"not null;size:255"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type LoginRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}
