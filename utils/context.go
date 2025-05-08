package utils

import (
	"github.com/gofiber/fiber/v3"
)

func GetUserID(c fiber.Ctx) (uint, error) {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return 0, fiber.NewError(fiber.StatusUnauthorized, "user_id not found or invalid")
	}
	return userID, nil
}
