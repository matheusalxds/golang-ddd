package interfaces

import (
	http "go-fx-project/src/internal/shared"
	"go-fx-project/src/internal/user/application"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService application.UserService
}

func NewUserHandler(userService application.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	type request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return http.BadRequest(c, "Invalid request")
	}

	newUser, err := h.userService.CreateUser(req.Name, req.Email)

	if err != nil {
		return http.BadRequest(c, "Could not create the user")
	}

	return http.Ok(c, newUser)
}
