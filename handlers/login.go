package handlers

import (
	"context"
	"time"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"api-rest/config"
	"api-rest/utils"
	"api-rest/models"
)

func Register(c *fiber.Ctx) error {
    user := new(models.User)

    if err := c.BodyParser(user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No se pudo parsear el body"})
    }

	// Validar campos requeridos
	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email y contraseña son obligatorios",
		})
	}

    user.ID = primitive.NewObjectID()
	// Hashear la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "No se pudo encriptar la contraseña",
		})
	}
	user.Password = string(hash)	

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel() 

	_, err = config.UsersCol.InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error al insertar usuario: " + err.Error(),
		})
	}

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Usuario registrado"})
}

func Login(c *fiber.Ctx) error {
// Parsear body con email y password
var req struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
if err := c.BodyParser(&req); err != nil {
    return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error al parsear datos"})
}

req.Email = strings.TrimSpace(req.Email)
req.Password = strings.TrimSpace(req.Password)

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

var user models.User
err := config.UsersCol.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
if err != nil {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciales inválidas"})
}

// Compara la contraseña usando bcrypt
err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
if err != nil {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Contraseña incorrecta"})
}

token, err := utils.CrearToken(user.ID, user.Nombre) // o user.Email
if err != nil {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "error": "No se pudo generar el token",
    })
}

return c.JSON(fiber.Map{
    "message": "Usuario logueado",
    "token": token,
})
}