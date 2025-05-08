package repository

import (
	"time"

	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"gorm.io/gorm"
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

func (tr *TransactionRepository) Delete(id uint) error {
	return tr.db.Delete(&entity.Transaction{}, id).Error
}

func (tr *TransactionRepository) FindByUser(userID uint) ([]entity.Transaction, error) {
	var Transactions []entity.Transaction
	if err := tr.db.Where("user_id = ?", userID).Find(&Transactions).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return Transactions, nil
}

func (tr *TransactionRepository) FindByTodo(todoID uint) ([]entity.Transaction, error) {
	var Transactions []entity.Transaction
	if err := tr.db.Where("todo_id = ?", todoID).Find(&Transactions).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return Transactions, nil
}

func (tr *TransactionRepository) FindByUserAndDateRange(userID uint, start, end time.Time) ([]entity.Transaction, error) {
	var Transactions []entity.Transaction
	if err := tr.db.Where("user_id = ? AND transaction_date BETWEEN ? AND ?", userID, start, end).Find(&Transactions).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return Transactions, nil
}

func (tr *TransactionRepository) FindByUserAndType(userID uint, transactionType string) ([]entity.Transaction, error) {
	var Transactions []entity.Transaction
	if err := tr.db.Where("user_id = ? AND type = ?", userID, transactionType).Find(&Transactions).Error; err != nil {
		return []entity.Transaction{}, err
	}
	return Transactions, nil
}


func (tr *TransactionRepository) TotalExpenseAmountByUserAndMonth(userID uint, year int, month int) (float32, error) {
	var expense_amount float32
	start := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)
	if err := tr.db.Model(&entity.Transaction{}).Select("user_id, sum(amount) as amount").Group("user_id").Having("user_id = ? AND type = 'expense' AND transaction_date BETWEEN ? AND ? ", userID, start, end).Find(&expense_amount).Error; err != nil {
		return 0, err
	}
	return expense_amount, nil
}

func (tr *TransactionRepository) TotalExpenseAmountByUserAndDateRange(userID uint, start, end time.Time) (float32, error) {
	var amount float32
	if err := tr.db.Model(&entity.Transaction{}).Select("user_id, sum(amount) as amount").Group("user_id").Having("user_id = ? AND type = 'expense' AND transaction_date BETWEEN ? AND ?", userID, start, end).Find(&amount).Error; err != nil {
		return 0, err
	}
	return amount, nil
}

func (tr *TransactionRepository) TotalReceiveAmountByUserAndMonth(userID uint, year int, month int) (float32, error) {
	var receive_amount float32
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0).Add(-time.Nanosecond)
	if err := tr.db.Model(&entity.Transaction{}).Select("user_id, sum(amount) as amount").Group("user_id").Having("user_id = ? AND type = 'receive' AND transaction_date BETWEEN ? AND ? ", userID, start, end).Find(&receive_amount).Error; err != nil {
		return 0, err
	}
	return receive_amount, nil
}

func (tr *TransactionRepository) TotalReceiveAmountByUserAndDateRange(userID uint, start, end time.Time) (float32, error) {
	var amount float32
	if err := tr.db.Model(&entity.Transaction{}).Select("user_id, sum(amount) as amount").Group("user_id").Having("user_id = ? AND type = 'receive' AND transaction_date BETWEEN ? AND ?", userID, start, end).Find(&amount).Error; err != nil {
		return 0, err
	}
	return amount, nil
}


func (tr *TransactionRepository) TotalAmountByUserAndMonth(userID uint, year int, month int) (float32, error) {
	expense, err := tr.TotalExpenseAmountByUserAndMonth(userID, year, month)
	if err != nil {
		return 0, err
	}
	receive, err := tr.TotalReceiveAmountByUserAndMonth(userID, year, month)
	if err != nil {
		return 0, err
	}
	return expense - receive, nil 
}

