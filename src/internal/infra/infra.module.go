package infra

import (
	"go-fx-project/src/internal/infra/db"
	"go-fx-project/src/internal/infra/env"
	idGenerator "go-fx-project/src/internal/infra/id-generator"
	"go-fx-project/src/internal/infra/logger"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(db.ConnectDatabase, idGenerator.NewIdGenerator, env.NewEnv, logger.NewLogger),
)
