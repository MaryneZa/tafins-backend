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
	app.Use("/todo", TodoRoutes(db))
	app.Use("/category", CategoryRoute(db))
	app.Use("/transaction",TransactionRoutes(db))
	app.Use("/daily_budget", DailyBudgetRoutes(db))
	app.Use("/monthly_budget", MonthlyBudgetRoutes(db))

	return app
}