package entity

import (
	"time"
	"gorm.io/gorm"
)

type Transaction struct{
	gorm.Model
	Type			string		`json:"type"`
	Title			string		`json:"title"`
	Amount			float32		`json:"amount"`
	TransactionDate *time.Time	`json:"transaction_date"`

	TodoID          uint		`json:"todo_id"`
	UserID			uint		`json:"user_id"`

	// Todo            Todo        `gorm:"foreignKey:TodoID"`
}