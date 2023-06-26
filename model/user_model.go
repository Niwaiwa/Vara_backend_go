package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Password     string    `json:"password" gorm:"not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	Nickname     string    `json:"nickname" gorm:"not null"`
	Avatar       string    `json:"avatar"`
	Header       string    `json:"header"`
	Description  string    `json:"description"`
	Created_at   time.Time `json:"created_at" gorm:"<-:create"`
	Updated_at   time.Time `json:"updated_at" gorm:"<-:update"`
	Last_login   time.Time `json:"last_login"`
	Locale       string    `json:"locale"`
	Is_active    bool      `json:"is_active"`
	Is_staff     bool      `json:"is_staff"`
	Is_superuser bool      `json:"is_superuser"`
}
