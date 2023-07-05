package main

import (
	"context"
	"errors"
	"fmt"
	"myapp/api/controller"
	"myapp/configs"
	"myapp/db"
	"myapp/domain"
	"myapp/repository"
	"myapp/usecase"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func setRoutes(e *echo.Echo, config *configs.Config, logger *zap.Logger, db *pgxpool.Pool) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	ur := repository.NewUserRepository(db, logger)
	lc := controller.NewLoginController(usecase.NewLoginUsecase(ur, logger), logger, db, *config)
	sc := controller.NewSignupController(usecase.NewSignupUsecase(ur, logger), logger, db, *config)
	rtc := controller.NewRefreshTokenController(usecase.NewRefreshTokenUsecase(ur, logger), logger, db, *config)
	e.POST("/login", lc.Login)
	e.POST("/signup", sc.Signup)
	e.POST("/refresh-token", rtc.RefreshToken)

	protectedRouter := e.Group("api")

	protectedRouter.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("test"),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JwtCustomClaims)
		},
		// SigningKey: []byte(config.SecretKey),
	}))

	protectedRouter.GET("/test", func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		c.Logger().Info(token)
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(*domain.JwtCustomClaims) // by default claims is of type `jwt.MapClaims`
		c.Logger().Info(claims)

		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	})

}

func main() {
	config := configs.InitConfigs()
	log.Info(config.Debug)
	// log.Info(config)

	e := echo.New()
	e.Debug = true
	// e.Logger.SetLevel(log.INFO)

	// use custom logger to inject logger into context
	logger := configs.InitLogger()
	defer logger.Sync()

	// connect to database and get database version to inject into context
	db, dbVersion := db.DbConnect(context.Background(), logger, config)
	defer db.Close()
	logger.Info(fmt.Sprintf("Connected to database. Version: %s", dbVersion))

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.NewString()
		},
	}))
	e.Use(configs.RequestLoggerMiddleware(logger))
	e.Use(middleware.Recover())

	setRoutes(e, config, logger, db)

	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
