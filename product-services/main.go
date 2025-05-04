package main

import "github.com/gofiber/fiber/v2"

func main() {
	getRespond()
}

func getRespond() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/users", func(r *fiber.Ctx) error {
		return r.JSON(fiber.Map{
			"data": []string{"Budi", "Rina", "Doni"},
		})
	})

	app.Listen(":8000")
}
