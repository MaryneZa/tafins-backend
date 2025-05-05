package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct{
	gorm.Model
	Name         string  		`json:"name"`
	Email        string  		`json:"email" gorm:"unique" validate:"required,email"`     
	Password     string  		`json:"password" validate:"required,min=8"`
	Age          uint8   		`json:"age" validate:"gte=0,lte=120"`
	Birthday     time.Time 		`json:"birthday"`
	Todos		 []Todo	 		`json:"todos" gorm:"foreignKey:UserID;references:ID"`
	Categories	 []Category 	`json:"categories" gorm:"foreignKey:UserID;references:ID"`
	Transactions []Transaction 	`json:"transactions" gorm:"foreignKey:UserID;references:ID"`
	MoneyGoals   []MoneyGoal	`json:"money_goals" gorm:"foreignKey:UserID;references:ID"`
	Budgets   	 []Budget		`json:"budgets" gorm:"foreignKey:UserID;references:ID"`
}

// Name         string  `json:"name" gorm:"default:null;not null"`