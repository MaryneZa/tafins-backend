package entity

import (
	"gorm.io/gorm"
)

type MonthlyBudget struct {
	gorm.Model
	Year  int `json:"year"`
	Month int `json:"month"` // 1-12

	LimitAmount float32 `json:"limit_amount"`

	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"` // preload user info alongside a todo
}
