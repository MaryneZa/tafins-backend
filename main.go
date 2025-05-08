package main

import (
	"log"

	"github.com/MaryneZa/tafins-backend/database"
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/routes"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.SetupDB()

	db.AutoMigrate(&entity.User{}, &entity.Todo{}, &entity.Transaction{}, &entity.MonthlyBudget{}, &entity.Category{}, &entity.DailyBudget{})

	app := routes.SetupRouter(db)

	app.Listen(":8090")
}
