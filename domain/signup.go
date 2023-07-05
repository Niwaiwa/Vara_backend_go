package domain

import (
	"go.uber.org/zap"
)

type SignupRequest struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(logger *zap.Logger, user *User) error
	GetUserByEmail(logger *zap.Logger, email string) (*User, error)
	GetUserByUsername(logger *zap.Logger, username string) (*User, error)
	CreateAccessToken(user *User, secret string, expire int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expire int) (refreshToken string, err error)
}
