package handler

import (
	"time"

	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type HttpUserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewHttpUserHandler(userUseCase usecase.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: userUseCase}
}

var validate = validator.New()

func (u *HttpUserHandler) SignUpHandler(c fiber.Ctx) error {
	newUser := new(entity.User)
	if err := c.Bind().Body(newUser); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := validate.Struct(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := u.userUseCase.SignUp(*newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("unable to create user " + err.Error())
	}

	return c.JSON(fiber.Map{"message": "sign up success !"})
}

func (u *HttpUserHandler) LogInHandler(c fiber.Ctx) error {
	loginUser := new(entity.User)
	if err := c.Bind().Body(loginUser); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := validate.Struct(loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	token, err := u.userUseCase.LogIn(*loginUser)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": token,
		"message":      "login successfully !",
	})
}
