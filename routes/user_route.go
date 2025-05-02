package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB) *fiber.App {
	app := fiber.New()

	app.Get("/test-auth", func(c fiber.Ctx) error {
		userID := c.Locals("user_id")
		return c.SendString(fmt.Sprintf("Hello, World! user_id : %d", userID))
	})

	return app
}