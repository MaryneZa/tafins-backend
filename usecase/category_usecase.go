package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type CategoryUseCase interface {
	CreateCategory(c entity.Category) error
	GetAllCategoryByUserID(userID uint) ([]entity.Category, error)
	// GetCategoryByID(categoryID uint) error
}

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) CategoryUseCase {
	return &CategoryService{repo: repo}
}

func (cs *CategoryService) CreateCategory(c entity.Category) error {
	return cs.repo.Create(c)
}

func (cs *CategoryService) GetAllCategoryByUserID(userID uint) ([]entity.Category, error) {
	return cs.repo.FindByUser(userID)
}

// func (cs *CategoryService) GetCategoryByID(id uint) error {
// 	return cs.repo.GetCategoryByID(id)
// }
