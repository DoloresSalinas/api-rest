package main

import (
	"log"

	"github.com/gofiber/fiber/v2" 
    "api-rest/routes"
    "api-rest/config" 
)

func main() {
	app := fiber.New() 

	config.ConectarMongo()
	
	routes.SetupRoutes(app)
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}