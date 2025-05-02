package entity

import (
	"time"
	"gorm.io/gorm"
)

type Budget struct {
	gorm.Model
	StartDate		time.Time		`json:"start_date"`
	EndDate			time.Time		`json:"end_date"`

	LimitAmount		float32			`json:"limit_amount"`

	UserID			uint			`json:"user_id"`
	User          	User			`gorm:"foreignKey:UserID"` // preload user info alongside a todo
}