package routes

import (
	"github.com/MaryneZa/tafins-backend/interface/handler"
	"github.com/MaryneZa/tafins-backend/interface/repository"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func MonthlyBudgetRoutes(db *gorm.DB) *fiber.App {
	app := fiber.New()

	monthlyBudgetRepository := repository.NewMonthlyBudgetRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	monthlyBudgetService := usecase.NewMonthlyBudgetService(monthlyBudgetRepository, transactionRepository)
	monthlyBudgetHandler := handler.NewHttpMonthlyBudgetHandler(monthlyBudgetService)

	app.Post("/create", monthlyBudgetHandler.CreateMonthlyBudgetHandler)
	app.Put("/update", monthlyBudgetHandler.UpdateMonthlyBudgetHandler)
	app.Get("/get", monthlyBudgetHandler.GetMonthlyBudgetHandler)
	app.Delete("/delete", monthlyBudgetHandler.DeleteMonthlyBudgetHandler)
	app.Get("/get-year", monthlyBudgetHandler.ListBudgetsForYearHandler)
	app.Get("/get-month-ramaining", monthlyBudgetHandler.GetRemainingMonthlyBudgetHandler)
	app.Get("/get-annual-budget", monthlyBudgetHandler.GetAnnualBudgetTotalHandler)

	return app
}
