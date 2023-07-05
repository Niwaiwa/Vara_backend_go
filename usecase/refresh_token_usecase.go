package usecase

import (
	"myapp/domain"
	"myapp/internal/tokenutil"

	"go.uber.org/zap"
)

type refreshTokenUsecase struct {
	userRepository domain.UserRepository
	logger         *zap.Logger
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, logger *zap.Logger) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(logger *zap.Logger, id string) (*domain.User, error) {
	return rtu.userRepository.GetByID(logger, id)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(logger *zap.Logger, user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user.ID.String(), secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(logger *zap.Logger, user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user.ID.String(), secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(logger *zap.Logger, requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
