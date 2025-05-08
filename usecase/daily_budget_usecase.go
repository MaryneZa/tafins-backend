package usecase

import (
	"time"
	"errors"
	"github.com/MaryneZa/tafins-backend/entity"
	"gorm.io/gorm"
)

type DailyBudgetUsecase interface {
	CreateBudget(dailyBudget entity.DailyBudget) error
	UpdateBudget(userID uint, date time.Time, amount float32) error
	GetBudgetByDate(userID uint, date time.Time) (entity.DailyBudget, error)
	DeleteBudget(userID uint, date time.Time) error
	ListBudgets(userID uint, startDate, endDate time.Time) ([]entity.DailyBudget, error)
	IsOverBudget(userID uint, date time.Time) (bool, error)
	GetTotalLimit(userID uint, startDate, endDate time.Time) (float32, error)
	// GenerateDailyBudgetsForMonth(userID uint, defaultLimit float32, year int, month int) error
}

type DailyBudgetService struct {
	repo DailyBudgetRepository
}

func NewDailyBudgetService(repo DailyBudgetRepository) DailyBudgetUsecase {
	return &DailyBudgetService{repo: repo}
}

func (ds *DailyBudgetService) CreateBudget(dailyBudget entity.DailyBudget) error {
	isExist, err := ds.repo.Exists(dailyBudget.UserID, dailyBudget.Date)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if isExist {
		return errors.New("budget already exists for this date")
	}

	return ds.repo.Create(dailyBudget)
}

func (ds *DailyBudgetService) UpdateBudget(userID uint, date time.Time, amount float32) error{
	return ds.repo.Update(userID, date, amount)
}

func (ds *DailyBudgetService) GetBudgetByDate(userID uint, date time.Time) (entity.DailyBudget, error){
	return ds.repo.FindByUserAndDate(userID, date)
}

func (ds *DailyBudgetService) DeleteBudget(userID uint, date time.Time) error{
	return ds.repo.Delete(userID, date)
}

func (ds *DailyBudgetService) ListBudgets(userID uint, startDate, endDate time.Time) ([]entity.DailyBudget, error) {
	return ds.repo.ListByDateRange(userID, startDate, endDate)
}

func (ds *DailyBudgetService) IsOverBudget(userID uint, date time.Time) (bool, error){
	return ds.repo.Exists(userID, date)
}

func (ds *DailyBudgetService) GetTotalLimit(userID uint, startDate, endDate time.Time) (float32, error) {
	return ds.repo.SumLimitByDateRange(userID, startDate, endDate)
}