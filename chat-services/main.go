package main

import (
	// Import controller yang baru
	authmodels "github.com/zackyfachrur/be-exalna-export/auth-services/models"
	"github.com/zackyfachrur/be-exalna-export/chat-services/internal/controllers"
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
	// -----------------------------------------------------
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_HOST_MYSQL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.ChatLog{}, &authmodels.User{}); err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	gClient, err := gemini.NewClient()
	if err != nil {
		log.Fatalf("Failed to initialize Gemini client: %v", err)
	}
	defer gClient.Close()

	// -----------------------------------------------------
	chatController := controllers.NewChatController(db, gClient)

	app.Post("/chat", chatController.HandlePostChat)
	app.Get("/chat/:userId", chatController.HandleGetUserChats)
	// -----------------------------------------------------

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}
