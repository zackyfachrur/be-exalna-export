package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/zackyfachrur/be-exalna-export/auth-services/config"
	"github.com/zackyfachrur/be-exalna-export/auth-services/controllers"
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
