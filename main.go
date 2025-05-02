package main

import (
	
	"log"
	"github.com/joho/godotenv"
	"github.com/MaryneZa/tafins-backend/database"
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/routes"

)


func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.SetupDB()

	db.AutoMigrate(&entity.User{}, &entity.Todo{}, &entity.Transaction{}, &entity.Budget{}, &entity.Category{}, &entity.MoneyGoal{})

	app := routes.SetupRouter(db)

	app.Listen(":8090")
}
