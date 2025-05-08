package usecase

import (
	"time"
	"github.com/MaryneZa/tafins-backend/entity"
)

type DailyBudgetRepository interface {
	Create(dailyBudget entity.DailyBudget) error
	Update(userID uint, date time.Time, amount float32) error
	FindByUserAndDate(userID uint, date time.Time) (entity.DailyBudget, error)

	Exists(userID uint, date time.Time) (bool, error)
	Delete(userID uint, date time.Time) error
	ListByDateRange(userID uint, startDate, endDate time.Time) ([]entity.DailyBudget, error)
	SumLimitByDateRange(userID uint, startDate, endDate time.Time) (float32, error)
}