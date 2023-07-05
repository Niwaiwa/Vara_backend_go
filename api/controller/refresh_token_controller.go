package controller

import (
	"myapp/configs"
	"myapp/domain"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type RefreshTokenController struct {
	refreshTokenUsecase domain.RefreshTokenUsecase
	logger              *zap.Logger
	database            *pgxpool.Pool
	config              configs.Config
}

func NewRefreshTokenController(refreshTokenUsecase domain.RefreshTokenUsecase, logger *zap.Logger, database *pgxpool.Pool, config configs.Config) *RefreshTokenController {
	return &RefreshTokenController{
		refreshTokenUsecase: refreshTokenUsecase,
		logger:              logger,
		database:            database,
		config:              config,
	}
}

func (rtc *RefreshTokenController) RefreshToken(c echo.Context) error {
	var request domain.RefreshTokenRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	id, err := rtc.refreshTokenUsecase.ExtractIDFromToken(rtc.logger, request.RefreshToken, rtc.config.SecretKey)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	user, err := rtc.refreshTokenUsecase.GetUserByID(rtc.logger, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	accessToken, err := rtc.refreshTokenUsecase.CreateAccessToken(rtc.logger, user, rtc.config.SecretKey, int(rtc.config.AccessTokenExpiryHour))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := rtc.refreshTokenUsecase.CreateRefreshToken(rtc.logger, user, rtc.config.SecretKey, int(rtc.config.RefreshTokenExpiryHour))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, refreshTokenResponse)
}
