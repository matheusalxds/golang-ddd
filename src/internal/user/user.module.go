package user

import (
	"go-fx-project/src/internal/user/application"
	"go-fx-project/src/internal/user/infra"
	interfaces "go-fx-project/src/internal/user/interface"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		infra.NewUserRepo,
		application.NewUserService,
		interfaces.NewUserHandler,
	),
	fx.Invoke(interfaces.RegisterUserRoutes),
)
