package repository

import (
	"gorm.io/gorm"
	"github.com/MaryneZa/tafins-backend/usecase"
	"github.com/MaryneZa/tafins-backend/entity"


)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) usecase.CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) Create(c entity.Category) error {
	return cr.db.Create(&c).Error
}

func (cr *CategoryRepository) Get(id uint) (entity.Category, error) {
	var category entity.Category
	result := cr.db.Where("id = ?", id).Find(&category)
	if result.Error!= nil {
		return entity.Category{}, result.Error
	}
	return category, nil
}

func (cr *CategoryRepository) GetAllByUserID(user_id uint) ([]entity.Category, error) {
	var categories []entity.Category
	if err := cr.db.Where("user_id = ?", user_id).Find(&categories).Error ; err != nil {
		return []entity.Category{}, err
	}
	return categories, nil
}
