package infra

import (
	"go-fx-project/src/internal/infra/db"
	idGenerator "go-fx-project/src/internal/infra/id-generator"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(db.ConnectDatabase, idGenerator.NewIdGenerator),
)
