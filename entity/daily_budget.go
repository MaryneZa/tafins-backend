package entity

import (
	"time"
	"gorm.io/gorm"
)

type DailyBudget struct {
	gorm.Model
    Date         time.Time `json:"date"`         // e.g., 2025-05-06
    LimitAmount  float32   `json:"limit_amount"` // e.g., 100.00
    UserID       uint      `json:"user_id"`
}
