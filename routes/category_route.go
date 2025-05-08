package routes

import (
	"github.com/MaryneZa/tafins-backend/interface/handler"
	"github.com/MaryneZa/tafins-backend/interface/repository"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func CategoryRoute(db *gorm.DB) *fiber.App {
	app := fiber.New()

	categoryRepo := repository.NewCategoryRepository(db)
	categorService := usecase.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewHttpCategoryHandler(categorService)

	app.Post("/create", categoryHandler.CreateCategoryHandler)
	app.Get("/get-all-mine", categoryHandler.GetAllCategoryHandler)

	return app
}
