package usecase

import (
	"time"

	"github.com/MaryneZa/tafins-backend/entity"
)

type TransactionRepository interface {
	Create(transaction entity.Transaction) error
	Delete(id uint) error
	FindByUser(userID uint) ([]entity.Transaction, error)
	FindByTodo(todoID uint) ([]entity.Transaction, error)

	FindByUserAndDateRange(userID uint, start, end time.Time) ([]entity.Transaction, error)
	FindByUserAndType(userID uint, transactionType string) ([]entity.Transaction, error)

	TotalExpenseAmountByUserAndMonth(userID uint, year int, month int) (float32, error)
	TotalExpenseAmountByUserAndDateRange(userID uint, start, end time.Time) (float32, error)
	TotalReceiveAmountByUserAndMonth(userID uint, year int, month int) (float32, error)
	TotalReceiveAmountByUserAndDateRange(userID uint, start, end time.Time) (float32, error)
	TotalAmountByUserAndMonth(userID uint, year int, month int) (float32, error)


}
