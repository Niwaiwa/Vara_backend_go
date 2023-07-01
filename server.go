package main

import (
	"context"
	"fmt"
	"myapp/configs"
	"myapp/db"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

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

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

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
