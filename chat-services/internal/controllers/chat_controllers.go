// internal/controllers/chat_controller.go
package controllers

import (
	authmodels "github.com/zackyfachrur/be-exalna-export/auth-services/models"
	"github.com/zackyfachrur/be-exalna-export/chat-services/internal/gemini"
	"github.com/zackyfachrur/be-exalna-export/chat-services/internal/models"
	"strconv" // Perlu untuk konversi userId dari string ke int

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

// ----------------------------
type ChatController struct {
	DB           *gorm.DB
	GeminiClient *gemini.Client
}

// ----------------------------------------------------
func NewChatController(db *gorm.DB, geminiClient *gemini.Client) *ChatController {
	return &ChatController{
		DB:           db,
		GeminiClient: geminiClient,
	}
}

// ----------------------------------------------------
func (cc *ChatController) HandlePostChat(c *fiber.Ctx) error {
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

	var user authmodels.User
	if err := cc.DB.First(&user, body.UserID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "User not found",
		})
	}

	response, err := cc.GeminiClient.GetWebsiteServices(c.Context(), body.Prompt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	chatLog := models.ChatLog{
		UserID:   body.UserID,
		Keyword:  body.Keyword,
		Prompt:   body.Prompt,
		Response: string(response),
	}

	if err := cc.DB.Create(&chatLog).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to save chat log",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    response,
	})
}

// ----------------------------------------------------
func (cc *ChatController) HandleGetUserChats(c *fiber.Ctx) error {
	userIDStr := c.Params("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid user ID format",
		})
	}

	var logs []models.ChatLog

	if err := cc.DB.Where("user_id = ?", userID).Find(&logs).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to fetch chat logs",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    logs,
	})
}
