package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"api-rest/config"
	"api-rest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
 

func GetTasks(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.TasksCol.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(tasks)
}


func CreateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input models.TaskInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No se pudo parsear el body"})
	}

	// Parsear fechas string -> time.Time
	fechaInicio, err := time.Parse("02/01/2006", input.FechaInicio)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha_inicio inválido"})
	}

	fechaDeadline, err := time.Parse("02/01/2006", input.FechaDeadline)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha_deadline inválido"})
	}

	// Parsear UserID string a ObjectID
	userID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "UserID inválido"})
	}

	task := models.Task{ 
		Titulo:        input.Titulo,
		Descripcion:   input.Descripcion,
		FechaInicio:   fechaInicio,
		FechaDeadline: fechaDeadline,
		UserID:        userID,
	}
    
    task.ID = primitive.NewObjectID()

	_, err = config.TasksCol.InsertOne(ctx, task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "No se pudo guardar la tarea"})
	}
	
    return c.Status(201).JSON(task)
}


func GetTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("ID inválido")
	}

	var task models.Task
	err = config.TasksCol.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return c.Status(404).SendString("Tarea no encontrada")
	}

	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("ID inválido")
	}

	var input models.TaskInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	updateFields := bson.M{}

	if input.Titulo != "" {
		updateFields["titulo"] = input.Titulo
	}
	if input.Descripcion != "" {
		updateFields["descripcion"] = input.Descripcion
	}
	if input.FechaInicio != "" {
		fechaInicio, err := time.Parse("02/01/2006", input.FechaInicio)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha_inicio inválido"})
		}
		updateFields["fecha_inicio"] = fechaInicio
	}
	if input.FechaDeadline != "" {
		fechaDeadline, err := time.Parse("02/01/2006", input.FechaDeadline)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Formato de fecha_deadline inválido"})
		}
		updateFields["fecha_deadline"] = fechaDeadline
	}
	if input.UserID != "" {
		userID, err := primitive.ObjectIDFromHex(input.UserID)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "UserID inválido"})
		}
		updateFields["user_id"] = userID
	}

	if len(updateFields) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No se proporcionaron campos para actualizar"})
	}

	update := bson.M{"$set": updateFields}

	_, err = config.TasksCol.UpdateByID(ctx, objID, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "Tarea actualizada correctamente",
		"id":      objID.Hex(),
	})
}


func DeleteTask(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(400).SendString("ID inválido")
	}

	_, err = config.TasksCol.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
