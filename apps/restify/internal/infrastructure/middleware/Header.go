package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireXUserHeader(c *fiber.Ctx) error {
	if c.Get("x-user") == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Missing x-user header")
	}
	return c.Next()
}
