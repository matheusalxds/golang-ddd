package logger

import (
	"go-fx-project/src/internal/infra/env"

	"go.uber.org/zap"
)

func NewLogger(env env.EnvLoader) *zap.Logger {
	var cfg zap.Config
	goEnv := env.LoadEnv().GoEnv

	if goEnv == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	logger, _ := cfg.Build()

	return logger
}
