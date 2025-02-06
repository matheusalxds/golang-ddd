package main

import (
	"context"
	"fmt"
	"go-fx-project/src/internal/infra"
	"go-fx-project/src/internal/infra/env"
	"go-fx-project/src/internal/user"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		fx.Provide(fiber.New),
		infra.Module,
		user.Module,
		fx.Invoke(startServer),
	)
	app.Run()
}

func startServer(lc fx.Lifecycle, app *fiber.App, env env.EnvLoader, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := env.LoadEnv().Port
			logger.Info("Server is running on http://localhost:" + port)
			go func() {
				if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping server...")
			return app.Shutdown()
		},
	})
}
