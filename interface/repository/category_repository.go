package repository

import (
	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"gorm.io/gorm"
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
	if result.Error != nil {
		return entity.Category{}, result.Error
	}
	return category, nil
}

func (cr *CategoryRepository) FindByUser(userID uint) ([]entity.Category, error) {
	var categories []entity.Category
	if err := cr.db.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return []entity.Category{}, err
	}
	return categories, nil
}
