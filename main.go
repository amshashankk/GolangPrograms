package main

import (
	"github.com/amshashankk/database"
	"github.com/amshashankk/routes"

	//	"github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	/* app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	})) */

	routes.Setup(app)

	app.Listen(":8000")
}