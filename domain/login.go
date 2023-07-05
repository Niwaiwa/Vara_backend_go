package domain

import (
	"go.uber.org/zap"
)

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetUserByEmail(logger *zap.Logger, email string) (*User, error)
	GetUserByUsername(logger *zap.Logger, username string) (*User, error)
	CreateAccessToken(user *User, secret string, expire int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expire int) (refreshToken string, err error)
}
