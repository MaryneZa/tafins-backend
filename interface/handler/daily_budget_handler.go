package handler

import (
	"time"
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
) 

type HttpDailyBudgetHandler struct {
	dailyBudgetUseCase usecase.DailyBudgetUsecase
}

func NewHttpDailyBudgetHandler(dailyBudgetUseCase usecase.DailyBudgetUsecase) *HttpDailyBudgetHandler {
	return &HttpDailyBudgetHandler{dailyBudgetUseCase: dailyBudgetUseCase}
}

type DailyBudgetDateRangeInput struct {
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}
type DailyBudgetInput struct {
	Date time.Time `json:"date"`
	Amount float32 `json:"amount"`
}


func (dh *HttpDailyBudgetHandler) CreateDailyBudgetHandler(c fiber.Ctx) error {
	dailyBudget := new(entity.DailyBudget)
	if err := c.Bind().Body(&dailyBudget); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	dailyBudget.UserID = userID
	if err := dh.dailyBudgetUseCase.CreateBudget(*dailyBudget); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create daily budget !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Create budget successfully !!"})
}

func (dh *HttpDailyBudgetHandler) DeleteDailyBudgetHandler(c fiber.Ctx) error {
	date := new(DailyBudgetInput)

	if err := c.Bind().Body(date); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	if err := dh.dailyBudgetUseCase.DeleteBudget(userID, date.Date); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot delete daily budget !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Creat daily budget successfully !!"})

}

func (dh *HttpDailyBudgetHandler) GetDailyBudgetByDateHandler(c fiber.Ctx) error {
	date := new(DailyBudgetInput)


	if err := c.Bind().Body(date); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	dailyBudget, err := dh.dailyBudgetUseCase.GetBudgetByDate(userID, date.Date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"daily_budget": dailyBudget,
	})
}

func (dh *HttpDailyBudgetHandler) GetTotalLimitDailyBudgetByDateRangeHandler(c fiber.Ctx) error {
	date := new(DailyBudgetDateRangeInput)

	if err := c.Bind().Body(date); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	amount, err := dh.dailyBudgetUseCase.GetTotalLimit(userID, date.StartDate, date.EndDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"amount": amount,
	})
}

func (dh *HttpDailyBudgetHandler) GetListDailyBudgetsByDateRangeHandler(c fiber.Ctx) error {
	date := new(DailyBudgetDateRangeInput)

	if err := c.Bind().Body(date); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	dailyBudgets, err := dh.dailyBudgetUseCase.ListBudgets(userID, date.StartDate, date.EndDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"daily_budgets": dailyBudgets,
	})
}

func (dh *HttpDailyBudgetHandler) UpdateDailyBudgetHandler(c fiber.Ctx) error {
	data := new(DailyBudgetInput)

	if err := c.Bind().Body(data); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	if err := dh.dailyBudgetUseCase.UpdateBudget(userID, data.Date, data.Amount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update daily budget !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Update daily budget successfully !!"})
}
