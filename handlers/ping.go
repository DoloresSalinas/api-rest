package handlers

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
    return c.SendString("Pong")
}
