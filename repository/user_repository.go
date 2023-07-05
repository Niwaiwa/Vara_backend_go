package repository

import (
	"context"
	"database/sql"
	"errors"
	"myapp/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var ErrAccountNotFound = errors.New("account not found")

type userRepository struct {
	database *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool, logger *zap.Logger) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(logger *zap.Logger, user *domain.User) error {
	query := `INSERT INTO users (id, username, password, email, nickname, avatar, header, description, created_at, updated_at, last_login, locale, is_active, is_staff, is_superuser) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW(), NOW(), $9, $10, $11, $12)`
	_, err := ur.database.Exec(context.Background(), query, user.ID, user.Username, user.Password, user.Email, user.Nickname, user.Avatar, user.Header, user.Description, user.Locale, user.Is_active, user.Is_staff, user.Is_superuser)
	if err != nil {
		logger.Error("Error creating user account.", zap.Error(err))
		return err
	}
	return nil
}

func (ur *userRepository) GetByEmail(logger *zap.Logger, emailParam string) (*domain.User, error) {
	var id uuid.UUID
	var username pgtype.Text
	var password pgtype.Text
	var email pgtype.Text
	var nickname pgtype.Text
	var avatar pgtype.Text
	var header pgtype.Text
	var description pgtype.Text
	var created_at pgtype.Timestamp
	var updated_at pgtype.Timestamp
	var last_login pgtype.Timestamp
	var locale pgtype.Text
	var is_active pgtype.Bool
	var is_staff pgtype.Bool
	var is_superuser pgtype.Bool

	query := `SELECT id, username, password, email, nickname, avatar, header, description, created_at, 
		updated_at, last_login, locale, is_active, is_staff, is_superuser 
		FROM users 
		WHERE email = $1`

	if err := ur.database.QueryRow(context.Background(), query, emailParam).Scan(&id, &username, &password, &email, &nickname, &avatar, &header, &description, &created_at, &updated_at, &last_login, &locale, &is_active, &is_staff, &is_superuser); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		logger.Error("Error retrieving user account.", zap.Error(err))
		return nil, err
	}

	return &domain.User{
		ID:           id,
		Username:     username.String,
		Password:     password.String,
		Email:        email.String,
		Nickname:     nickname.String,
		Avatar:       avatar.String,
		Header:       header.String,
		Description:  description.String,
		Created_at:   created_at.Time,
		Updated_at:   updated_at.Time,
		Last_login:   last_login.Time,
		Locale:       locale.String,
		Is_active:    is_active.Bool,
		Is_staff:     is_staff.Bool,
		Is_superuser: is_superuser.Bool,
	}, nil
}

func (ur *userRepository) GetByUsername(logger *zap.Logger, usernameParam string) (*domain.User, error) {
	var id uuid.UUID
	var username pgtype.Text
	var password pgtype.Text
	var email pgtype.Text
	var nickname pgtype.Text
	var avatar pgtype.Text
	var header pgtype.Text
	var description pgtype.Text
	var created_at pgtype.Timestamp
	var updated_at pgtype.Timestamp
	var last_login pgtype.Timestamp
	var locale pgtype.Text
	var is_active pgtype.Bool
	var is_staff pgtype.Bool
	var is_superuser pgtype.Bool

	query := `SELECT id, username, password, email, nickname, avatar, header, description, created_at, 
		updated_at, last_login, locale, is_active, is_staff, is_superuser 
		FROM users 
		WHERE username = $1`

	if err := ur.database.QueryRow(context.Background(), query, usernameParam).Scan(&id, &username, &password, &email, &nickname, &avatar, &header, &description, &created_at, &updated_at, &last_login, &locale, &is_active, &is_staff, &is_superuser); err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		logger.Error("Error retrieving user account.", zap.Error(err))
		return nil, err
	}

	return &domain.User{
		ID:           id,
		Username:     username.String,
		Password:     password.String,
		Email:        email.String,
		Nickname:     nickname.String,
		Avatar:       avatar.String,
		Header:       header.String,
		Description:  description.String,
		Created_at:   created_at.Time,
		Updated_at:   updated_at.Time,
		Last_login:   last_login.Time,
		Locale:       locale.String,
		Is_active:    is_active.Bool,
		Is_staff:     is_staff.Bool,
		Is_superuser: is_superuser.Bool,
	}, nil
}

func (ur *userRepository) GetByID(logger *zap.Logger, userID string) (*domain.User, error) {
	var id uuid.UUID
	var username pgtype.Text
	var password pgtype.Text
	var email pgtype.Text
	var nickname pgtype.Text
	var avatar pgtype.Text
	var header pgtype.Text
	var description pgtype.Text
	var created_at pgtype.Timestamp
	var updated_at pgtype.Timestamp
	var last_login pgtype.Timestamp
	var locale pgtype.Text
	var is_active pgtype.Bool
	var is_staff pgtype.Bool
	var is_superuser pgtype.Bool

	query := `SELECT id, username, password, email, nickname, avatar, header, description, created_at, 
		updated_at, last_login, locale, is_active, is_staff, is_superuser 
		FROM users 
		WHERE id = $1`

	if err := ur.database.QueryRow(context.Background(), query, userID).Scan(&id, &username, &password, &email, &nickname, &avatar, &header, &description, &created_at, &updated_at, &last_login, &locale, &is_active, &is_staff, &is_superuser); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrAccountNotFound
		}
		logger.Error("Error retrieving user account.", zap.Error(err))
		return nil, err
	}

	return &domain.User{
		ID:           id,
		Username:     username.String,
		Password:     password.String,
		Email:        email.String,
		Nickname:     nickname.String,
		Avatar:       avatar.String,
		Header:       header.String,
		Description:  description.String,
		Created_at:   created_at.Time,
		Updated_at:   updated_at.Time,
		Last_login:   last_login.Time,
		Locale:       locale.String,
		Is_active:    is_active.Bool,
		Is_staff:     is_staff.Bool,
		Is_superuser: is_superuser.Bool,
	}, nil
}
