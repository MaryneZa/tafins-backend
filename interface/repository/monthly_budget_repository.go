package repository


import (
	"gorm.io/gorm"
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
)

type MonthlyBudgetRepository struct {
	db *gorm.DB
}

func NewMonthlyBudgetRepository(db *gorm.DB) usecase.MonthlyBudgetRepository {
	return &MonthlyBudgetRepository{db: db}
}


func (mr *MonthlyBudgetRepository) Create(budget entity.MonthlyBudget) error {
	return mr.db.Create(&budget).Error
}
func (mr *MonthlyBudgetRepository) Update(userID uint, year int, month int, newAmount float32) error {
	return mr.db.Model(&entity.MonthlyBudget{}).Where("user_id = ? AND year = ? AND month = ?", userID, year, month).Update("limit_amount", newAmount).Error
}
func (mr *MonthlyBudgetRepository) FindByUserAndMonth(userID uint, year int, month int) (entity.MonthlyBudget, error) {
	var monthlyBudget entity.MonthlyBudget
	if err := mr.db.Where("user_id = ? AND year = ? AND month = ?", userID, year, month).Find(&monthlyBudget).Error; err != nil {
		return entity.MonthlyBudget{}, err
	}
	return monthlyBudget, nil
}

func (mr *MonthlyBudgetRepository) FindLimitValueByUserAndMonth(userID uint, year int, month int) (float32, error) {
	var amount float32
	if err := mr.db.Table("monthly_budgets").
			Select("limit_amount").
			Where("user_id = ? AND year = ? AND month = ?", userID, year, month).
			Scan(&amount).Error; err != nil {
		return 0, err
	}
	return amount, nil
}


func (mr *MonthlyBudgetRepository) Delete(userID uint, year int, month int) error {
	if err := mr.db.Where("user_id = ? AND year = ? AND month = ?", userID, year, month).Delete(&entity.MonthlyBudget{}).Error; err != nil {
		return err
	} 
	return nil
}
func (mr *MonthlyBudgetRepository) Exists(userID uint, year int, month int) (bool, error) {
	monthlyBudget, err := mr.FindByUserAndMonth(userID, year, month)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err // Some other DB error
	}
	return monthlyBudget.ID != 0, nil
}
func (mr *MonthlyBudgetRepository) ListByUserAndYear(userID uint, year int) ([]entity.MonthlyBudget, error) {
	var monthlyBudgets []entity.MonthlyBudget

	if err := mr.db.Where("user_id = ? AND year = ?", userID, year).
		Order("month DESC").
		Find(&monthlyBudgets).
		Error; err != nil {
			return []entity.MonthlyBudget{}, err
		}
	
	if len(monthlyBudgets) == 0 {
		return []entity.MonthlyBudget{}, nil
	}

	return monthlyBudgets, nil
}
func (mr *MonthlyBudgetRepository) SumAnnualBudget(userID uint, year int) (float32, error) {
	var amount float32
	if err := mr.db.Model(&entity.MonthlyBudget{}).
	Select("COALESCE(SUM(limit_amount), 0)").
	Where("user_id = ? AND year = ?", userID, year).Find(&amount).Error; err != nil {
		return 0, err
	}
	return amount, nil
}