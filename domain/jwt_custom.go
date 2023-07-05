package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
