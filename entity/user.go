package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string          `json:"name"`
	Email          string          `json:"email" gorm:"unique" validate:"required,email"`
	Password       string          `json:"password" validate:"required,min=8"`
	Age            uint8           `json:"age" validate:"gte=0,lte=120"`
	Birthday       time.Time       `json:"birthday"`
	Todos          []Todo          `json:"todos" gorm:"foreignKey:UserID;references:ID"`
	Categories     []Category      `json:"categories" gorm:"foreignKey:UserID;references:ID"`
	Transactions   []Transaction   `json:"transactions" gorm:"foreignKey:UserID;references:ID"`
	DailyGoals     []DailyBudget    `json:"daily_budgets" gorm:"foreignKey:UserID;references:ID"`
	MonthlyBudgets []MonthlyBudget `json:"monthly_budgets" gorm:"foreignKey:UserID;references:ID"`
}

// Name         string  `json:"name" gorm:"default:null;not null"`
