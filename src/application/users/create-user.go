package applications

import (
	usecases "go-fx-project/src/usecases/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService usecases.UserService
}

func NewUserHandler(userService usecases.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	type request struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	newUser, err := h.userService.CreateUser(req.Name, req.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create the user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfuly",
		"user":    newUser,
	})
}
