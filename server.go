package main

import (
	"context"
	"myapp/configs"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := configs.InitConfigs()
	log.Info(config.Debug)

	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.INFO)

	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.TimeKey = "time"
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	logger, err := zapConfig.Build()
	defer logger.Sync()
	if err != nil {
		panic(err)
	}

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.NewString()
		},
	}))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogError:     true,
		LogRequestID: true,
		LogLatency:   true,
		LogMethod:    true,
		HandleError:  true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.Info("request",
					zap.String("request_id", v.RequestID),
					zap.String("method", v.Method),
					zap.String("URI", v.URI),
					zap.Int64("latency", v.Latency.Nanoseconds()),
					zap.Int("status", v.Status),
				)
			} else {
				logger.Error("request error",
					zap.String("request_id", v.RequestID),
					zap.String("method", v.Method),
					zap.String("URI", v.URI),
					zap.Int64("latency", v.Latency.Nanoseconds()),
					zap.Int("status", v.Status),
					zap.Error(v.Error),
				)
			}
			return nil
		},
	}))
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
