package httpHelper

import "github.com/gofiber/fiber/v2"

func ErrorResponse(c *fiber.Ctx, statusCode int, msg string) error {
	return c.Status(statusCode).JSON(fiber.Map{"statusCode": statusCode, msg: msg})
}

func BadRequest(c *fiber.Ctx, msg string) error {
	return ErrorResponse(c, fiber.StatusBadRequest, msg)
}

func Ok(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"statusCode": fiber.StatusOK, "data": data})
}
