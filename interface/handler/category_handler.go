package handler

import (
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/MaryneZa/tafins-backend/utils"
	"github.com/gofiber/fiber/v3"
)

type HttpCategoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
}

func NewHttpCategoryHandler(c usecase.CategoryUseCase) *HttpCategoryHandler {
	return &HttpCategoryHandler{categoryUseCase: c}
}

func (hc *HttpCategoryHandler) CreateCategoryHandler(c fiber.Ctx) error {
	category := new(entity.Category)

	userID, err := utils.GetUserID(c)
	if err != nil {
		return err
	}

	category.UserID = userID

	if err := c.Bind().Body(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input => " + err.Error()})
	}
	if err := validate.Struct(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input => " + err.Error()})
	}
	if err := hc.categoryUseCase.CreateCategory(*category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create category !!"})
	}
	return c.Status(fiber.StatusOK).JSON((fiber.Map{"massage": "Create category successfully !"}))
}

func (hc *HttpCategoryHandler) GetAllCategoryHandler(c fiber.Ctx) error {
	userID, err := utils.GetUserID(c)
	if err != nil {
		return err
	}

	categories, err := hc.categoryUseCase.GetAllCategoryByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retreive all category !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"categories": categories,
	})
}
