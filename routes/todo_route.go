package routes

import (
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v3"
	"github.com/MaryneZa/tafins-backend/interface/handler"
	"github.com/MaryneZa/tafins-backend/interface/repository"
	"github.com/MaryneZa/tafins-backend/usecase"
)

func TodoRoutes(db *gorm.DB) *fiber.App {

	app := fiber.New()

	todoRepo := repository.NewTodoRepository(db)
	todoService := usecase.NewTodoService(todoRepo)
	todoHandler := handler.NewHttpTodoHandler(todoService)

	app.Post("/create", todoHandler.CreateTodoHandler)
	app.Get("/get-all", todoHandler.GetAllTodosHandler)

	return app

}