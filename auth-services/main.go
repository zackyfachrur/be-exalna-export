package main

import (
	"auth-services/config"
	"auth-services/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Connect to database
	config.ConnectDatabase()

	// Middleware
	app.Use(cors.New())

	// Routes
	app.Post("/login", controllers.LoginUser)
	app.Post("/register", controllers.RegisterUser)

	// Start server
	if err := app.Listen(":3001"); err != nil {
		panic(err)
	}
}
