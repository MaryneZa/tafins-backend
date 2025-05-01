package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct{
	gorm.Model
	Name         string  `json:"name" validate:"required"`
	Email        string  `json:"email" gorm:"unique" validate:"required,email"`     
	Password     string  `json:"password" validate:"required,min=8"`
	Age          uint8   `json:"age" validate:"gte=0,lte=120"`
	Birthday     *time.Time 
}

// Name         string  `json:"name" gorm:"default:null;not null"`