package main

import (
	"context"
	"fmt"
	applications "go-fx-project/src/application/users"
	"go-fx-project/src/infra/db"
	repo "go-fx-project/src/infra/db/repo/users"
	idGenerator "go-fx-project/src/infra/id-generator"
	usecases "go-fx-project/src/usecases/user"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			db.ConnectDatabase,
			NewFiberApp,
			repo.NewUserRepo,
			applications.NewUserHandler,
			idGenerator.NewIdGenerator,
			usecases.NewUserService,
		),
		fx.Invoke(registerRoutes),
		fx.Invoke(startServer),
	)

	app.Run()
}

func NewFiberApp() *fiber.App {
	return fiber.New()
}

func registerRoutes(app *fiber.App, userHandler *applications.UserHandler) {
	app.Post("/users", userHandler.CreateUser)
}

func startServer(lc fx.Lifecycle, app *fiber.App) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := 3000
			fmt.Printf("Server is running on http://localhost:%d\n", port)
			go func() {
				if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping server...")
			return app.Shutdown()
		},
	})
}
