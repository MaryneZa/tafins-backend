package handler

import (
	"strconv"
	"time"
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/gofiber/fiber/v3"
)

type HttpTodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func NewHttpTodoHandler(todoUseCase usecase.TodoUseCase) *HttpTodoHandler {
	return &HttpTodoHandler{todoUseCase: todoUseCase}
}

type CreateTodoInput struct {
	Title      string    `json:"title" validate:"required"`
	Date       time.Time `json:"date" validate:"required"`
	CategoryID *uint      `json:"category_id"`
}

type UpdateTodoInput struct {
	ID 		   uint      `json:"id" validate:"required"`
	Title      string    `json:"title"`
	Date       time.Time `json:"date"`
}

func (th *HttpTodoHandler) CreateTodoHandler(c fiber.Ctx) error {
	todo_detail := new(CreateTodoInput)
	if err := c.Bind().Body(todo_detail); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := validate.Struct(todo_detail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input => " + err.Error()})
	}
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}
	
	todo := entity.Todo{
		Title: todo_detail.Title,
		Date: todo_detail.Date,
		CategoryID: todo_detail.CategoryID,
		UserID: userID,
	}

	if err := th.todoUseCase.CreateTodo(todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot create todo !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Create todo successfully !!"})
}

func (th *HttpTodoHandler) UpdateTodoHandler(c fiber.Ctx) error {
	todo := new(entity.Todo)
	if err := c.Bind().Body(todo); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// userID, ok := c.Locals("user_id").(uint)
	// if !ok {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": "user_id not found or invalid",
	// 	})
	// }
	// todo.UserID = userID

	if err := th.todoUseCase.UpdateTodo(*todo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot update todo !!" + err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Update todo successfully !!"})

}

func (th *HttpTodoHandler) GetAllTodoByUserIDHandler(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found or invalid",
		})
	}

	todos, err := th.todoUseCase.GetAllTodosByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retreive all todo !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todos" : todos, 
	})
}

func (th *HttpTodoHandler) GetTodoByIDHandler(c fiber.Ctx) error {
	idStr := c.Params("id")

	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid ID format",
		})
	}

	id := uint(idUint64)

	todo, err := th.todoUseCase.GetTodo(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retreive todo !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todo" : todo, 
	})

}

func (th *HttpTodoHandler) GetAllTodosHandler(c fiber.Ctx) error {
	todos, err := th.todoUseCase.GetAllTodos(); 
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot retreive all todo !!"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todos" : todos, 
	})
}


