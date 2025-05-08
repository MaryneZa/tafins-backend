package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type CategoryRepository interface {
	Create(todo entity.Category) error
	Get(categoryID uint) (entity.Category, error)
	FindByUser(userID uint) ([]entity.Category, error)
}
