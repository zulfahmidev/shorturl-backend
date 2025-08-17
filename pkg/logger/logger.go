package logger

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func InitLogger() (err error) {
	env := os.Getenv("APP_ENV")

	if env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
    defer logger.Sync()
	return
}

func Default() *zap.Logger {
	return logger
}

type LogRequest struct {
	requestID string
	logger    *zap.Logger
	tag       string
}

func NewRequest(tag string) LogRequest {
	requestID := uuid.New().String()
	logger = logger.With(zap.String("requestID", requestID))
	return LogRequest{
		requestID: requestID,
		logger:    logger,
		tag:       tag,
	}
}

func (l *LogRequest) Err(subtag string, message string) {
	if l.requestID != "" {
		tag := fmt.Sprintf("%s-%s", l.tag, subtag)
		l.logger.Error(message, zap.String("tag", tag))
	}
}

func (l *LogRequest) Info(subtag string, message string) {
	tag := fmt.Sprintf("%s-%s", l.tag, subtag)
	l.logger.Info(message, zap.String("tag", tag))
}
