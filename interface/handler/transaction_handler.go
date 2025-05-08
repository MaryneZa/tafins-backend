package handler

import (
	"strconv"
	"time"

	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
)

type HttpTransactionHandler struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewHttpTransactionHandler(transactionUseCase usecase.TransactionUseCase) *HttpTransactionHandler {
	return &HttpTransactionHandler{transactionUseCase: transactionUseCase}
}

// type CreateTransactionInput struct {
// 	Type            string     `json:"type" validate:"required"` // "expense" or "receive"
// 	Title           string     `json:"title" validate:"required"`
// 	Amount          float32    `json:"amount" validate:"required,gt=0"`
// 	TransactionDate *time.Time `json:"transaction_date"`
// 	TodoID          uint       `json:"todo_id"`
// 	UserID          uint       `json:"user_id"`
// }

type DateRangeInput struct {
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}

type MonthYearInput struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}


func (th *HttpTransactionHandler) CreateTransactionHandler(c fiber.Ctx) error {
	t := new(entity.Transaction)
	if err := c.Bind().Body(&t); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	t.UserID = userID
	if err := th.transactionUseCase.CreateTransaction(*t); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Create transaction successfully !"})
}

func (th *HttpTransactionHandler) GetAllTransactionByTodoIDHandler(c fiber.Ctx) error {
	idStr := c.Params("id")

	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid ID format",
		})
	}

	id := uint(idUint64)

	transactions, err := th.transactionUseCase.GetAllTransactionByTodoID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transactions": transactions,
	})
}

func (th *HttpTransactionHandler) GetAllTransactionByUserIDHandler(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}

	transactions, err := th.transactionUseCase.GetAllTransactionByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transactions": transactions,
	})

}

func (th *HttpTransactionHandler) FindByUserAndDateRangeHandler(c fiber.Ctx) error {
	date := new(DateRangeInput)
	
	if err := c.Bind().Body(&date); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}

	transactions, err := th.transactionUseCase.FindByUserAndDateRange(userID, date.StartDate, date.EndDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transactions": transactions,
	})

}

func (th *HttpTransactionHandler) FindByUserAndTypeHandler(c fiber.Ctx) error {
	var transactionType string
	
	if err := c.Bind().Body(&transactionType); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}

	transactions, err := th.transactionUseCase.FindByUserAndType(userID, transactionType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transactions": transactions,
	})

}

func (th *HttpTransactionHandler) TotalExpenseAmountByUserAndDateRangeHandler(c fiber.Ctx) error {
	date := new(DateRangeInput)
	
	if err := c.Bind().Body(&date); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}

	amount, err := th.transactionUseCase.TotalExpenseAmountByUserAndDateRange(userID, date.StartDate, date.EndDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"amount": amount,
	})

}

func (th *HttpTransactionHandler) TotalExpenseAmountByUserAndMonth(c fiber.Ctx) error {
	input := new(MonthYearInput)
	
	if err := c.Bind().Body(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}

	amount, err := th.transactionUseCase.TotalExpenseAmountByUserAndMonth(userID, input.Year, input.Month)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"amount": amount,
	})

}

func (th *HttpTransactionHandler) TotalReceiveAmountByUserAndDateRangeHandler(c fiber.Ctx) error {
	date := new(DateRangeInput)
	
	if err := c.Bind().Body(&date); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}

	amount, err := th.transactionUseCase.TotalReceiveAmountByUserAndDateRange(userID, date.StartDate, date.EndDate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"amount": amount,
	})

}

func (th *HttpTransactionHandler) TotalReceiveAmountByUserAndMonth(c fiber.Ctx) error {
	input := new(MonthYearInput)
	
	if err := c.Bind().Body(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}

	amount, err := th.transactionUseCase.TotalReceiveAmountByUserAndMonth(userID, input.Year, input.Month)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"amount": amount,
	})

}



