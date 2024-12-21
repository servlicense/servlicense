package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	scope := c.Locals("scopes")
	return c.JSON(fiber.Map{
		"success": true,
		"message": scope,
	})
}
