package routes

import (
	"github.com/MaryneZa/tafins-backend/interface/handler"
	"github.com/MaryneZa/tafins-backend/interface/repository"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func AuthRoutes(db *gorm.DB) *fiber.App {
	app := fiber.New()

	userRepo := repository.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handler.NewHttpUserHandler(userService)

	app.Post("/signup", userHandler.SignUpHandler)
	app.Post("/login", userHandler.LogInHandler)

	return app
}
