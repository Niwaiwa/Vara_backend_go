package usecase

import (
	"myapp/domain"
	"myapp/internal/tokenutil"

	"go.uber.org/zap"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	logger         *zap.Logger
}

func NewSignupUsecase(userRepository domain.UserRepository, logger *zap.Logger) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (su *signupUsecase) Create(logger *zap.Logger, user *domain.User) error {
	return su.userRepository.Create(logger, user)
}

func (su *signupUsecase) GetUserByEmail(logger *zap.Logger, email string) (*domain.User, error) {
	return su.userRepository.GetByEmail(logger, email)
}

func (su *signupUsecase) GetUserByUsername(logger *zap.Logger, username string) (*domain.User, error) {
	return su.userRepository.GetByUsername(logger, username)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expire int) (string, error) {
	return tokenutil.CreateAccessToken(user.ID.String(), secret, expire)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expire int) (string, error) {
	return tokenutil.CreateRefreshToken(user.ID.String(), secret, expire)
}
