package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type TransactionRepository interface {
	Create(t entity.Transaction) error
	GetAllByUserID(user_id uint) ([]entity.Transaction, error)
	GetAllByTodoID(todo_id uint) ([]entity.Transaction, error)
}