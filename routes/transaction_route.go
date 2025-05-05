package routes

import (
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v3"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/MaryneZa/tafins-backend/interface/handler"
	"github.com/MaryneZa/tafins-backend/interface/repository"
)

func TransactionRoutes(db *gorm.DB) *fiber.App {
	app := fiber.New()

	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := usecase.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewHttpTransactionHandler(transactionService)

	app.Post("/create", transactionHandler.CreateTransactionHandler)
	app.Get("/get-all-mine", transactionHandler.GetAllTransactionByUserIDHandler)
	app.Get("/get-todo/:id", transactionHandler.GetAllTransactionByTodoIDHandler)

	return app
}