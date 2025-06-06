package middleware

import (
	"strings"  
	"github.com/gofiber/fiber/v2"
	"api-rest/utils"
)

func AuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Token no proporcionado",
            })
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Formato de token inválido",
            })
        }

        tokenString := parts[1]
        token, err := utils.ValidarToken(tokenString)
        if err != nil || !token.Valid {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Token inválido",
            })
        }

        return c.Next()
    }
}
