package entity

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name		string		`json:"name" validate:"required"`
	UserID		uint		`json:"user_id" validate:"required"`

	Todos 		[]Todo 		`gorm:"foreignKey:CategoryID;references:ID"`
}