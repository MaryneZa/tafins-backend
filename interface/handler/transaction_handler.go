package handler

import (
	"strconv"
	"github.com/gofiber/fiber/v3"
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
)

type HttpTransactionHandler struct {
	transactionUseCase usecase.TransactionUseCase
}

func NewHttpTransactionHandler(transactionUseCase usecase.TransactionUseCase) *HttpTransactionHandler {
	return &HttpTransactionHandler{transactionUseCase: transactionUseCase}
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
		"transactions" : transactions,
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
		"transactions" : transactions,
	})

}

