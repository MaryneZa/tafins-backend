package handler

import (

	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
)

type HttpMonthlyBudgetHandler struct {
	monthlyBudgetUseCase usecase.MonthlyBudgetUseCase
}

func NewHttpMonthlyBudgetHandler(monthlyBudgetHandler usecase.MonthlyBudgetUseCase) *HttpMonthlyBudgetHandler {
	return &HttpMonthlyBudgetHandler{monthlyBudgetUseCase: monthlyBudgetHandler}
}

type MonthlyBudgetInput struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

func (mh *HttpMonthlyBudgetHandler) CreateMonthlyBudgetHandler(c fiber.Ctx) error {
	monthlyBudget := new(entity.MonthlyBudget)
	if err := c.Bind().Body(monthlyBudget); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	monthlyBudget.UserID = userID
	if err := mh.monthlyBudgetUseCase.CreateBudget(monthlyBudget.UserID, monthlyBudget.Year, monthlyBudget.Month, monthlyBudget.LimitAmount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create monthly budget !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Create budget successfully !!"})
}
func (mh *HttpMonthlyBudgetHandler) UpdateMonthlyBudgetHandler(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	monthlyBudget := new(entity.MonthlyBudget)
	if err := c.Bind().Body(monthlyBudget); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := mh.monthlyBudgetUseCase.UpdateBudget(userID, monthlyBudget.Year, monthlyBudget.Month, monthlyBudget.LimitAmount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create monthly budget !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Create budget successfully !!"})
}
func (mh *HttpMonthlyBudgetHandler) GetMonthlyBudgetHandler(c fiber.Ctx) error {
	monthlyBudgetInput := new(MonthlyBudgetInput)
	if err := c.Bind().Body(monthlyBudgetInput); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	monthlyBudget, err := mh.monthlyBudgetUseCase.GetMonthlyBudget(userID, monthlyBudgetInput.Year, monthlyBudgetInput.Month)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"monthly_budget": monthlyBudget,
	})
}
func (mh *HttpMonthlyBudgetHandler) DeleteMonthlyBudgetHandler(c fiber.Ctx) error {
	input := new(MonthlyBudgetInput)
	if err := c.Bind().Body(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	if err := mh.monthlyBudgetUseCase.DeleteMonthlyBudget(userID, input.Year, input.Month); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("delete monthly budget successfully !")
}
func (mh *HttpMonthlyBudgetHandler) ListBudgetsForYearHandler(c fiber.Ctx) error {
	input := new(MonthlyBudgetInput)
	if err := c.Bind().Body(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	monthlyBudgets, err := mh.monthlyBudgetUseCase.ListBudgetsForYear(userID, input.Year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"monthly_budgets": monthlyBudgets,
	})
}
func (mh *HttpMonthlyBudgetHandler) GetRemainingMonthlyBudgetHandler(c fiber.Ctx) error {
	input := new(MonthlyBudgetInput)
	if err := c.Bind().Body(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	amount, err := mh.monthlyBudgetUseCase.GetRemainingMonthlyBudget(userID, input.Year, input.Month)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"amount": amount,
	})
}
func (mh *HttpMonthlyBudgetHandler) GetAnnualBudgetTotalHandler(c fiber.Ctx) error {
	input := new(MonthlyBudgetInput)
	if err := c.Bind().Body(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	amount, err := mh.monthlyBudgetUseCase.GetAnnualBudgetTotal(userID, input.Year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"amount": amount,
	})
	
}

