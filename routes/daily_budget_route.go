package routes

import (
	"github.com/MaryneZa/tafins-backend/interface/handler"
	"github.com/MaryneZa/tafins-backend/interface/repository"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func DailyBudgetRoutes(db *gorm.DB) *fiber.App {
	app := fiber.New()

	dailyBudgetRepository := repository.NewDailyBudgetRepository(db)
	dailyBudgetService := usecase.NewDailyBudgetService(dailyBudgetRepository)
	dailyBudgetHandler := handler.NewHttpDailyBudgetHandler(dailyBudgetService)

	app.Post("/create", dailyBudgetHandler.CreateDailyBudgetHandler)

	app.Get("/get-date", dailyBudgetHandler.GetDailyBudgetByDateHandler)
	app.Get("/get-date-range", dailyBudgetHandler.GetListDailyBudgetsByDateRangeHandler)
	app.Get("/get-date-range/total", dailyBudgetHandler.GetTotalLimitDailyBudgetByDateRangeHandler)
	
	app.Put("/update", dailyBudgetHandler.UpdateDailyBudgetHandler)

	app.Delete("/delete", dailyBudgetHandler.DeleteDailyBudgetHandler)

	return app
}
