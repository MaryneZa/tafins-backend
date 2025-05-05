package repository

import (
	"gorm.io/gorm"
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"

)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) usecase.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) Create(t entity.Transaction) error {
	return tr.db.Create(&t).Error
}

func (tr *TransactionRepository) GetAllByUserID(user_id uint) ([]entity.Transaction, error) {
	var Transactions []entity.Transaction
	if err := tr.db.Where("user_id = ?", user_id).Find(&Transactions).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return Transactions, nil
}

func (tr *TransactionRepository) GetAllByTodoID(todo_id uint) ([]entity.Transaction, error) {
	var Transactions []entity.Transaction
	if err := tr.db.Where("todo_id = ?", todo_id).Find(&Transactions).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return Transactions, nil
}

