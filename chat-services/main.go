package main

import (
	authmodels "github.com/zackyfachrur/be-exalna-export/auth-services/models"
	"github.com/zackyfachrur/be-exalna-export/chat-services/internal/gemini"
	"github.com/zackyfachrur/be-exalna-export/chat-services/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"log"
	"os"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_HOST_MYSQL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate ChatLog dari internal dan User dari auth-services
	if err := db.AutoMigrate(&models.ChatLog{}, &authmodels.User{}); err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Init Gemini Client
	gClient, err := gemini.NewClient()
	if err != nil {
		log.Fatalf("Failed to initialize Gemini client: %v", err)
	}
	defer gClient.Close()

	// POST /chat
	app.Post("/chat", func(c *fiber.Ctx) error {
		var body gemini.GeminiRequest

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid request body",
			})
		}

		if body.Prompt == "" || body.UserID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Prompt and user_id are required",
			})
		}

		// Validasi user dari authmodels
		var user authmodels.User
		if err := db.First(&user, body.UserID).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "User not found",
			})
		}

		// Proses ke Gemini
		response, err := gClient.GetWebsiteServices(c.Context(), body.Prompt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		}

		// Simpan hasil ke database
		chatLog := models.ChatLog{
			UserID:   body.UserID,
			Keyword:  body.Keyword,
			Prompt:   body.Prompt,
			Response: string(response),
		}

		if err := db.Create(&chatLog).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   "Failed to save chat log",
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"data":    response,
		})
	})

	// GET /chat/:userId
	app.Get("/chat/:userId", func(c *fiber.Ctx) error {
		userID := c.Params("userId")
		var logs []models.ChatLog

		if err := db.Where("user_id = ?", userID).Find(&logs).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   "Failed to fetch chat logs",
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"data":    logs,
		})
	})

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
