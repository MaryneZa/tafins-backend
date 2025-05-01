package repository

import (
	"github.com/MaryneZa/tafins/entity"
	"github.com/MaryneZa/tafins/usecase"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) usecase.UserRepository{
	return &UserRepository{db: db}
}

func (ur *UserRepository) Save(user entity.User) error {
	return ur.db.Create(&user).Error
}

func (ur *UserRepository) Get(user entity.User) (entity.User, error) {
	var userDetail entity.User
	result := ur.db.Where("email = ?", user.Email).Find(&userDetail)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return userDetail, nil
}

func (ur *UserRepository) Find(email string) error {
	var userDetail entity.User
	result := ur.db.Where("email = ?", email).Find(&userDetail)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
