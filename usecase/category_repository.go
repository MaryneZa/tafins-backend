package usecase

import(
	"github.com/MaryneZa/tafins-backend/entity"

)

type CategoryRepository interface {
	Create(todo entity.Category) error
	Get(category_id uint) (entity.Category, error)
	GetAllByUserID(user_id uint) ([]entity.Category, error)
}