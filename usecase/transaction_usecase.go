package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type TransactionUseCase interface {
	CreateTransaction(t entity.Transaction) error
	GetAllTransactionByUserID(user_id uint) ([]entity.Transaction, error)
	GetAllTransactionByTodoID(todo_id uint) ([]entity.Transaction, error)
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
	return ts.repo.GetAllByTodoID(id)
}

func (ts *TransactionService) GetAllTransactionByUserID(id uint) ([]entity.Transaction, error) {
	return ts.repo.GetAllByUserID(id)
}