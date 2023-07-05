package domain

import (
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID           uuid.UUID `json:"-"`
	Username     string    `json:"username" binding:"required"`
	Password     string    `json:"password" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Header       string    `json:"header"`
	Description  string    `json:"description"`
	Created_at   time.Time `json:"-"`
	Updated_at   time.Time `json:"-"`
	Last_login   time.Time `json:"last_login"`
	Locale       string    `json:"locale" default:"en-US"`
	Is_active    bool      `json:"is_active"`
	Is_staff     bool      `json:"is_staff"`
	Is_superuser bool      `json:"is_superuser"`
}

type UserRepository interface {
	Create(logger *zap.Logger, user *User) error
	GetByEmail(logger *zap.Logger, email string) (*User, error)
	GetByUsername(logger *zap.Logger, username string) (*User, error)
	GetByID(logger *zap.Logger, id string) (*User, error)
}
