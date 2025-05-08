package repository

import (
	"errors"

	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/MaryneZa/tafins-backend/usecase"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) usecase.TodoRepository {
	return &TodoRepository{db: db}
}

func (tr *TodoRepository) Create(todo entity.Todo) error {
	return tr.db.Create(&todo).Error
}

func (tr *TodoRepository) Get(id uint) (entity.Todo, error) {
	var todo entity.Todo
	result := tr.db.Where("id = ?", id).Find(&todo)
	if result.Error != nil {
		return entity.Todo{}, result.Error
	}
	return todo, nil
}

func (tr *TodoRepository) FindAll() ([]entity.Todo, error) {
	var todos []entity.Todo
	if err := tr.db.Preload("User").
		Preload("Category").
		Preload("Transactions").
		Order("date DESC").
		Find(&todos).Error; err != nil {
		return []entity.Todo{}, errors.New("Error retrieving todos: " + err.Error())
	}
	return todos, nil
}

func (tr *TodoRepository) FindByUser(userID uint) ([]entity.Todo, error) {
	var todos []entity.Todo

	if err := tr.db.Where("user_id = ?", userID).
		Preload("User").
		Preload("Category").
		Preload("Transactions").
		Order("date DESC").
		Find(&todos).Error; err != nil {
		return []entity.Todo{}, errors.New("Error retrieving todos: " + err.Error())
	}

	if len(todos) == 0 {
		return []entity.Todo{}, nil
	}

	return todos, nil
}

func (tr *TodoRepository) Update(todo entity.Todo) error {
	result := tr.db.Model(&todo).Update("title", &todo.Title)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (tr *TodoRepository) Delete(id uint) error {
	result := tr.db.Delete(&entity.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
