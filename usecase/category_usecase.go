package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type CategoryUseCase interface {
	CreateCategory(c entity.Category) error
	GetAllCategoryByUserID(user_id uint) ([]entity.Category,error)
	GetCategoryByID(category_id uint) error
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

func (cs *CategoryService) GetAllCategoryByUserID(user_id uint) ([]entity.Category,error) {
	return cs.repo.GetAllByUserID(user_id)  
}

func (cs *CategoryService) GetCategoryByID(id uint) error {
	return cs.GetCategoryByID(id)
}

