package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type TodoRepository interface {
	Create(todo entity.Todo) error
	Get(todoID uint) (entity.Todo, error)
	FindAll() ([]entity.Todo, error)
	FindByUser(userID uint) ([]entity.Todo, error)
	Update(todo entity.Todo) error
	Delete(todoID uint) error
}
