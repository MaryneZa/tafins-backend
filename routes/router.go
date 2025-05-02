package routes

import (
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v3"
	"github.com/MaryneZa/tafins-backend/middleware"

)


func SetupRouter(db *gorm.DB) *fiber.App{
	app := fiber.New()

	app.Use("/auth", AuthRoutes(db))

	app.Use(middleware.AuthMiddleware)
	
	app.Use("/user", UserRoutes(db))

	return app
}