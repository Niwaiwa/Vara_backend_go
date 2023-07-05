package controller

import (
	"myapp/configs"
	"myapp/domain"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	loginUsecase domain.LoginUsecase
	logger       *zap.Logger
	database     *pgxpool.Pool
	config       configs.Config
}

func NewLoginController(loginUsecase domain.LoginUsecase, logger *zap.Logger, database *pgxpool.Pool, config configs.Config) *LoginController {
	return &LoginController{
		loginUsecase: loginUsecase,
		logger:       logger,
		database:     database,
		config:       config,
	}
}

func (lc *LoginController) Login(c echo.Context) error {
	var request domain.LoginRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	user, err := lc.loginUsecase.GetUserByUsername(lc.logger, request.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "User not found with the given username"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
	}

	accessToken, err := lc.loginUsecase.CreateAccessToken(user, lc.config.SecretKey, int(lc.config.AccessTokenExpiryHour))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := lc.loginUsecase.CreateRefreshToken(user, lc.config.SecretKey, int(lc.config.RefreshTokenExpiryHour))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, loginResponse)
}
