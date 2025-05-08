package usecase

import (
	"errors"
	"github.com/MaryneZa/tafins-backend/entity"
)

type MonthlyBudgetUseCase interface {
	CreateBudget(userID uint, year int, month int, limit float32) error
	UpdateBudget(userID uint, year int, month int, newLimit float32) error
	GetMonthlyBudget(userID uint, year int, month int) (entity.MonthlyBudget, error)
	DeleteMonthlyBudget(userID uint, year int, month int) error
	ListBudgetsForYear(userID uint, year int) ([]entity.MonthlyBudget, error)
	IsOverMonthlyBudget(userID uint, year int, month int) (bool, error)
	GetRemainingMonthlyBudget(userID uint, year int, month int) (float32, error)
	GetAnnualBudgetTotal(userID uint, year int) (float32, error)
}


type MonthlyBudgetService struct {
	monthlyRepo MonthlyBudgetRepository
	transactionRepo TransactionRepository
}

func NewMonthlyBudgetService(monthlyRepo MonthlyBudgetRepository, transactionRepo TransactionRepository) MonthlyBudgetUseCase {
	return &MonthlyBudgetService{monthlyRepo: monthlyRepo, transactionRepo: transactionRepo}
}

func (ms *MonthlyBudgetService) CreateBudget(userID uint, year int, month int, limit float32) error {
	isExist, err := ms.monthlyRepo.Exists(userID, year, month)
	if err != nil {
		return err
	}

	if isExist {
		return errors.New("budget already exists for this date")
	}
	data := entity.MonthlyBudget{
		UserID: userID,
		Year: year,
		Month: month,
		LimitAmount: limit,
	}
	return ms.monthlyRepo.Create(data)
}

func (ms *MonthlyBudgetService) UpdateBudget(userID uint, year int, month int, newLimit float32) error {
	return ms.monthlyRepo.Update(userID, year, month, newLimit)
}
func (ms *MonthlyBudgetService) GetMonthlyBudget(userID uint, year int, month int) (entity.MonthlyBudget, error) {
	return ms.monthlyRepo.FindByUserAndMonth(userID, year, month)
}
func (ms *MonthlyBudgetService) DeleteMonthlyBudget(userID uint, year int, month int) error {
	return ms.monthlyRepo.Delete(userID, year, month)
}
func (ms *MonthlyBudgetService) ListBudgetsForYear(userID uint, year int) ([]entity.MonthlyBudget, error) {
	return ms.monthlyRepo.ListByUserAndYear(userID, year)
}
func (ms *MonthlyBudgetService) IsOverMonthlyBudget(userID uint, year int, month int) (bool, error) {
	amount, err := ms.GetRemainingMonthlyBudget(userID, year, month)
	if err != nil {
		return false, err
	}
	return amount < 0, nil

}
func (ms *MonthlyBudgetService) GetRemainingMonthlyBudget(userID uint, year int, month int) (float32, error) {
	monthly_budget, err := ms.monthlyRepo.FindLimitValueByUserAndMonth(userID, year, month)
	if err != nil {
		return 0, err
	}
	total_daily, err := ms.transactionRepo.TotalAmountByUserAndMonth(userID, year, month)
	if err != nil {
		return 0, err
	}
	return monthly_budget - total_daily, nil
}
func (ms *MonthlyBudgetService) GetAnnualBudgetTotal(userID uint, year int) (float32, error) {
	return ms.monthlyRepo.SumAnnualBudget(userID, year)
}