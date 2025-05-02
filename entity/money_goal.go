package entity

import (
	"time"
	"gorm.io/gorm"
)

type MoneyGoal struct {
	gorm.Model
	TargetDate		time.Time		`json:"target_date"`
	LimitAmount		float32			`json:"limit_amount"`

	UserID			uint			`json:"user_id"`
	User          	User			`gorm:"foreignKey:UserID"` // preload user info alongside a todo
}