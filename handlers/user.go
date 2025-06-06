package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"api-rest/config"
	"api-rest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"golang.org/x/crypto/bcrypt"
)


func GetUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.UsersCol.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(users)
}


func GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("ID inválido")
	}

	var user models.User
	err = config.UsersCol.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return c.Status(404).SendString("Usuario no encontradao")
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("ID inválido")
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	updateFields := bson.M{}

	if user.Nombre != "" {
		updateFields["nombre"] = user.Nombre
	} 
	if user.App != "" {
		updateFields["app"] = user.App
	} 
	if user.Apm != "" {
		updateFields["apm"] = user.Apm
	} 
	if user.Email != "" {
		updateFields["email"] = user.Email
	} 
	if user.Password != "" {
		// Hashear la contraseña
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "No se pudo encriptar la contraseña",
			})
		}
		updateFields["password"] = string(hash)	
	} 
	if !user.FechaNacimiento.IsZero() {
		updateFields["fecha_nacimiento"] = user.FechaNacimiento
	}
	if user.PreguntaSecreta != "" {
		updateFields["pregunta_secreta"] = user.PreguntaSecreta
	} 
	if user.RespuestaSecreta != "" {
		updateFields["respuesta_secreta"] = user.RespuestaSecreta
	} 
	if len(updateFields) == 0 {
		return c.Status(400).SendString("No se enviaron campos para actualizar")
	} 

	update := bson.M{"$set": updateFields}

	_, err = config.UsersCol.UpdateByID(ctx, objID, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user.ID = objID
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("ID inválido")
	}

	_, err = config.UsersCol.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
