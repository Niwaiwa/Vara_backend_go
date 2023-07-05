package db

import (
	"context"
	"fmt"
	"myapp/configs"
	"net/url"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxuuid "github.com/vgarvardt/pgx-google-uuid/v5"
	"go.uber.org/zap"
)

func DbConnect(ctx context.Context, logger *zap.Logger, config *configs.Config) (*pgxpool.Pool, string) {
	rawURL := config.DBURL
	if !(strings.HasPrefix(rawURL, "postgresql://") || strings.HasPrefix(rawURL, "postgres://")) {
		rawURL = fmt.Sprintf("postgres://%s", rawURL)
	}
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		logger.Fatal("Bad database connection URL", zap.Error(err))
	}
	query := parsedURL.Query()
	var queryUpdated bool
	if len(query.Get("sslmode")) == 0 {
		query.Set("sslmode", "prefer")
		queryUpdated = true
	}
	if queryUpdated {
		parsedURL.RawQuery = query.Encode()
	}

	dbconfig, err := pgxpool.ParseConfig(parsedURL.String())
	if err != nil {
		logger.Fatal("Bad database connection URL", zap.Error(err))
	}
	dbconfig.MaxConnLifetime = time.Duration(config.DBConnMaxLifetimeMs) * time.Millisecond
	dbconfig.MaxConns = config.DBMaxOpenConns
	dbconfig.MinConns = config.DBMaxIdleConns
	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	logger.Debug("Complete database connection URL", zap.String("raw_url", parsedURL.String()))
	// dbPool, err := pgxpool.New(context.Background(), parsedURL.String())
	dbPool, err := pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		logger.Fatal("Error connecting to database", zap.Error(err))
	}

	var dbVersion string
	if err = dbPool.QueryRow(context.Background(), "SELECT version()").Scan(&dbVersion); err != nil {
		logger.Fatal("Error querying database version", zap.Error(err))
	}
	return dbPool, dbVersion
}
