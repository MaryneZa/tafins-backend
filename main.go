package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"strconv"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v3"
	"github.com/MaryneZa/tafins/entity"
	"github.com/MaryneZa/tafins/interface/repository"
	"github.com/MaryneZa/tafins/interface/handler"
	"github.com/MaryneZa/tafins/usecase"
	"github.com/MaryneZa/tafins/middleware"
)

func setupDB() *gorm.DB {
	
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB_NAME")
	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal("can not load db_port from .env")
	}
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
	host, user, password, dbname, port)

	log.Println(dsn)
	
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:   time.Second,   // Slow SQL threshold
		  LogLevel:        logger.Info, // Log level
		  Colorful:        true,          // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect to database "+ err.Error())
	}

	log.Println("connect to database successfully !!")

	return db
}

func main() {

	db := setupDB()

	db.AutoMigrate(&entity.User{})

	app := fiber.New()

	userRepo := repository.NewUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	userHandler := handler.NewHttpUserHandler(userService)

	app.Post("/signup", userHandler.SignUpHandler)
	app.Post("/login", userHandler.LogInHandler)

	app.Use(middleware.AuthMiddleware)

	app.Get("/test-auth", func(c fiber.Ctx) error {
		userID := c.Locals("user_id")
		return c.SendString(fmt.Sprintf("Hello, World! user_id : %d", userID))
	})
	

	app.Listen(":8090")
}