package domain

import "go.uber.org/zap"

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(logger *zap.Logger, id string) (*User, error)
	CreateAccessToken(logger *zap.Logger, user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(logger *zap.Logger, user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(logger *zap.Logger, requestToken string, secret string) (string, error)
}
