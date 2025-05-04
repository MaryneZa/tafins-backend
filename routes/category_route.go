package routes

import (
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v3"
	"github.com/MaryneZa/tafins-backend/interface/handler"
	"github.com/MaryneZa/tafins-backend/interface/repository"
	"github.com/MaryneZa/tafins-backend/usecase"
)

func CategoryRoute(db *gorm.DB) *fiber.App{
	app := fiber.New()

	categoryRepo := repository.NewCategoryRepository(db)
	categorService := usecase.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewHttpCategoryHandler(categorService)

	app.Post("/create", categoryHandler.CreateCategoryHandler)
	app.Get("/get-all-mine", categoryHandler.GetAllCategoryHandler)

	return app
} 