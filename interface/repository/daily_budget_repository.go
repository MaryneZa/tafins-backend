package repository


import (
	"errors"
	"time"
	"gorm.io/gorm"
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
)

type DailyBudgetRepository struct {
	db *gorm.DB
}

func NewDailyBudgetRepository(db *gorm.DB) usecase.DailyBudgetRepository {
	return &DailyBudgetRepository{db: db}
}


func (dr *DailyBudgetRepository) Create(dailyBudgets entity.DailyBudget) error {
	return dr.db.Create(&dailyBudgets).Error
}

func (dr *DailyBudgetRepository) Update(userID uint, date time.Time, amount float32) error {
	return dr.db.Model(&entity.DailyBudget{}).Where("user_id = ? AND date = ?", userID, date).Update("limit_amount", amount).Error
}


func (dr *DailyBudgetRepository) FindByUserAndDate(userID uint, date time.Time) (entity.DailyBudget, error) {
	var dailyBudget entity.DailyBudget
	err := dr.db.Where("user_id = ? AND date = ?", userID, date).First(&dailyBudget).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.DailyBudget{}, err
		}
		return entity.DailyBudget{}, err
	}

	return dailyBudget, nil
}


func (dr *DailyBudgetRepository) Exists(userID uint, date time.Time) (bool, error) {
	dailyBudget, err := dr.FindByUserAndDate(userID, date)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err // Some other DB error
	}
	return dailyBudget.ID != 0, nil
}

func (dr *DailyBudgetRepository) Delete(userID uint, date time.Time) error {
	if err := dr.db.Where("user_id = ? AND date = ?", userID, date).Delete(&entity.DailyBudget{}).Error; err != nil {
		return err
	} 
	return nil
}
func (dr *DailyBudgetRepository) ListByDateRange(userID uint, startDate, endDate time.Time) ([]entity.DailyBudget, error) {
	var dailyBudgets []entity.DailyBudget

	if err := dr.db.Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
		Order("date DESC").
		Find(&dailyBudgets).
		Error; err != nil {
			return []entity.DailyBudget{}, err
		}
	
	if len(dailyBudgets) == 0 {
		return []entity.DailyBudget{}, nil
	}

	return dailyBudgets, nil
}


func (dr *DailyBudgetRepository) SumLimitByDateRange(userID uint, startDate, endDate time.Time) (float32, error) {
	var amount float32
	if err := dr.db.Model(&entity.DailyBudget{}).
			Select("COALESCE(SUM(limit_amount), 0)").
			Where("user_id = ? AND date BETWEEN ? AND ?", userID, startDate, endDate).
			Find(&amount).Error; err != nil {
		return 0, err
	}
	return amount, nil
}

