package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.Logger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.TimeKey = "time"
	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

func RequestLoggerMiddleware(logger *zap.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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
	})
}
