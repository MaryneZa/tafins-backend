package entity

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	ID     			uint   			`json:"id" gorm:"primarykey"`
	Title			string			`json:"title"`
	Status			string			`json:"status"`
	Date			time.Time		`json:"date"`

	UserID			uint			`json:"user_id"`
	User          	User			`gorm:"foreignKey:UserID"` // preload user info alongside a todo

	CategoryID      *uint			`json:"category_id"`
	Category    	*Category		`gorm:"foreignKey:CategoryID"`

	Transactions	[]Transaction	`json:"transactions" gorm:"foreignKey:TodoID;references:ID"`
}