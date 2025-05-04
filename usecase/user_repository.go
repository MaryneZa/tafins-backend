package usecase

import "github.com/MaryneZa/tafins-backend/entity"

type UserRepository interface {
	Create(user entity.User) error
	Get(user entity.User) (entity.User, error)
	Find(email string) error
	// Delete(user entity.User) error
}
