package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type TodoRepository interface {
	Create(todo entity.Todo) error
	Get(todo_id uint) (entity.Todo, error)
	GetAll() ([]entity.Todo, error)
	GetAllByUserID(user_id uint) ([]entity.Todo, error)
	Update(todo entity.Todo) error
	Delete(todo_id uint) error
}
