package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type MonthlyBudgetRepository interface {
	Create(budget entity.MonthlyBudget) error
	Update(userID uint, year int, month int, newAmount float32) error
	FindByUserAndMonth(userID uint, year int, month int) (entity.MonthlyBudget, error)
	FindLimitValueByUserAndMonth(userID uint, year int, month int) (float32, error)
	Delete(userID uint, year int, month int) error
	Exists(userID uint, year int, month int) (bool, error)
	ListByUserAndYear(userID uint, year int) ([]entity.MonthlyBudget, error)
	SumAnnualBudget(userID uint, year int) (float32, error)
}
