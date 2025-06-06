package routes

import (
	"github.com/gofiber/fiber/v2"
	"api-rest/handlers"
	"api-rest/middleware"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/ping", handlers.Ping)
	app.Get("/users", handlers.GetUsers)
}

func Setup(app *fiber.App) {
	app.Get("/ping", handlers.Ping)

	api := app.Group("/api")

	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)
 
	user := api.Group("/users", middleware.AuthMiddleware())
	user.Get("/", handlers.GetUsers) 
    user.Get("/:id", handlers.GetUser)
    user.Put("/:id", handlers.UpdateUser)
    user.Delete("/:id", handlers.DeleteUser)

	task := api.Group("/tasks", middleware.AuthMiddleware())
	task.Get("/", handlers.GetTasks)
	task.Post("/", handlers.CreateTask)
    task.Get("/:id", handlers.GetTask)
    task.Put("/:id", handlers.UpdateTask)
    task.Delete("/:id", handlers.DeleteTask)
}
