package middleware

import "github.com/gofiber/fiber/v2"

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Log request details
		// You can customize this with your preferred logging format
		println("Request:", c.Method(), c.Path())
		return c.Next()
	}
}