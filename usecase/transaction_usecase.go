package usecase

import (
	"time"
	"github.com/MaryneZa/tafins-backend/entity"
)

type TransactionUseCase interface {
	CreateTransaction(t entity.Transaction) error
	GetAllTransactionByUserID(userID uint) ([]entity.Transaction, error)
	GetAllTransactionByTodoID(todoID uint) ([]entity.Transaction, error)
	FindByUserAndDateRange(userID uint, start, end time.Time) ([]entity.Transaction, error)
	FindByUserAndType(userID uint, transactionType string) ([]entity.Transaction, error)

	TotalExpenseAmountByUserAndMonth(userID uint, year int, month int) (float32, error)
	TotalExpenseAmountByUserAndDateRange(userID uint, start, end time.Time) (float32, error)
	TotalReceiveAmountByUserAndMonth(userID uint, year int, month int) (float32, error)
	TotalReceiveAmountByUserAndDateRange(userID uint, start, end time.Time) (float32, error)

	TotalAmountByUserAndMonth(userID uint, year int, month int) (float32, error)

}

type TransactionService struct {
	repo TransactionRepository
}

func NewTransactionService(repo TransactionRepository) TransactionUseCase {
	return &TransactionService{repo: repo}
}

func (ts *TransactionService) CreateTransaction(t entity.Transaction) error {
	return ts.repo.Create(t)
}

func (ts *TransactionService) GetAllTransactionByTodoID(id uint) ([]entity.Transaction, error) {
	return ts.repo.FindByTodo(id)
}

func (ts *TransactionService) GetAllTransactionByUserID(id uint) ([]entity.Transaction, error) {
	return ts.repo.FindByUser(id)
}

func (ts *TransactionService) FindByUserAndDateRange(userID uint, start, end time.Time) ([]entity.Transaction, error){
	return ts.repo.FindByUserAndDateRange(userID, start, end)
}

func (ts *TransactionService) FindByUserAndType(userID uint, transactionType string) ([]entity.Transaction, error){
	return ts.repo.FindByUserAndType(userID, transactionType) 
}

func (ts *TransactionService) TotalExpenseAmountByUserAndMonth(userID uint, year int, month int) (float32, error){
	return ts.repo.TotalExpenseAmountByUserAndMonth(userID, year, month)
}

func (ts *TransactionService) TotalExpenseAmountByUserAndDateRange(userID uint, start, end time.Time) (float32, error){
	return ts.repo.TotalExpenseAmountByUserAndDateRange(userID, start, end)
}

func (ts *TransactionService) TotalReceiveAmountByUserAndMonth(userID uint, year int, month int) (float32, error){
	return ts.repo.TotalReceiveAmountByUserAndMonth(userID, year, month)
}

func (ts *TransactionService) TotalReceiveAmountByUserAndDateRange(userID uint, start, end time.Time) (float32, error){
	return ts.repo.TotalReceiveAmountByUserAndDateRange(userID, start, end)
}

func (ts *TransactionService) TotalAmountByUserAndMonth(userID uint, year int, month int) (float32, error) {
	return ts.repo.TotalAmountByUserAndMonth(userID, year, month)
}

