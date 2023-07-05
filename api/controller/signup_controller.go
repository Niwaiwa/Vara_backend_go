package controller

import (
	"myapp/configs"
	"myapp/domain"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	logger        *zap.Logger
	database      *pgxpool.Pool
	config        configs.Config
}

func NewSignupController(signupUsecase domain.SignupUsecase, logger *zap.Logger, database *pgxpool.Pool, config configs.Config) *SignupController {
	return &SignupController{
		SignupUsecase: signupUsecase,
		logger:        logger,
		database:      database,
		config:        config,
	}
}

func (sc *SignupController) Signup(c echo.Context) error {
	var request domain.SignupRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	_, err = sc.SignupUsecase.GetUserByUsername(sc.logger, request.Username)
	if err == nil {
		return c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given username"})
	}

	_, err = sc.SignupUsecase.GetUserByEmail(sc.logger, request.Email)
	if err == nil {
		return c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       uuid.New(),
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	err = sc.SignupUsecase.Create(sc.logger, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.config.SecretKey, int(sc.config.AccessTokenExpiryHour))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.config.SecretKey, int(sc.config.RefreshTokenExpiryHour))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, signupResponse)
}
