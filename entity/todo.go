package entity

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title			string			`json:"title"`
	Status			string			`json:"status"`
	Date			time.Time		`json:"time"`

	UserID			uint			`json:"user_id"`
	User          	User			`gorm:"foreignKey:UserID"` // preload user info alongside a todo

	CategoryID      uint			`json:"category_id"`
	Category    	Category		`gorm:"foreignKey:CategoryID"`

	Transactions	[]Transaction	`json:"transactions" gorm:"foreignKey:TodoID;references:ID"`
}