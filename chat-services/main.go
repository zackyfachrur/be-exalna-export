package main

import (
	"log"
	"os"

	"bot-services/internal/gemini"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize Fiber
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Initialize Gemini Client
	gClient, err := gemini.NewClient()
	if err != nil {
		log.Fatalf("Failed to initialize Gemini client: %v", err)
	}
	defer gClient.Close()

	// Register routes
	app.Post("/chat", func(c *fiber.Ctx) error {
		var body gemini.GeminiRequest

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid request Body",
			})
		}

		if body.Prompt == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Prompt is Required",
			})
		}

		// Ganti dengan pemanggilan GenerateContent jika sesuai
		response, err := gClient.GetWebsiteServices(c.Context(), body.Prompt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}
		// log.Println("Prompt:", body.Prompt)
		// log.Println("Raw Gemini Response:", string(response))

		// Pastikan response sudah sesuai format
		return c.JSON(fiber.Map{
			"success": true,
			"data":    string(response), // Jika response dalam bentuk byte, convert ke string
		})
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}
	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
