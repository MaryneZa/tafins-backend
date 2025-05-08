package routes

import (
	"github.com/MaryneZa/tafins-backend/interface/handler"
	"github.com/MaryneZa/tafins-backend/interface/repository"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func TodoRoutes(db *gorm.DB) *fiber.App {

	app := fiber.New()

	todoRepo := repository.NewTodoRepository(db)
	todoService := usecase.NewTodoService(todoRepo)
	todoHandler := handler.NewHttpTodoHandler(todoService)

	app.Post("/create", todoHandler.CreateTodoHandler)
	app.Get("/get-all", todoHandler.GetAllTodosHandler)
	app.Get("/get-all-mine", todoHandler.GetAllTodoByUserIDHandler)
	app.Put("/update", todoHandler.UpdateTodoHandler)

	return app

}
