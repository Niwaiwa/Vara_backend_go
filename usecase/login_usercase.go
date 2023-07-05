package usecase

import (
	"myapp/domain"
	"myapp/internal/tokenutil"

	"go.uber.org/zap"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	logger         *zap.Logger
}

func NewLoginUsecase(userRepository domain.UserRepository, logger *zap.Logger) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (lu *loginUsecase) GetUserByEmail(logger *zap.Logger, email string) (*domain.User, error) {
	return lu.userRepository.GetByEmail(logger, email)
}

func (lu *loginUsecase) GetUserByUsername(logger *zap.Logger, username string) (*domain.User, error) {
	return lu.userRepository.GetByUsername(logger, username)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expire int) (string, error) {
	return tokenutil.CreateAccessToken(user.ID.String(), secret, expire)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expire int) (string, error) {
	return tokenutil.CreateRefreshToken(user.ID.String(), secret, expire)
}
