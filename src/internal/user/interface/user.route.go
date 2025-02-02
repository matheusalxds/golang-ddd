package interfaces

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(app *fiber.App, userHandler *UserHandler) {
	app.Post("/users", userHandler.CreateUser)
}
