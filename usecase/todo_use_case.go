package usecase

import (
	"github.com/MaryneZa/tafins-backend/entity"
)

type TodoUseCase interface {
	CreateTodo(todo entity.Todo) error
	GetTodo(id uint) (entity.Todo, error)
	GetAllTodos() ([]entity.Todo, error)
	GetAllTodosByUserID(userID uint) ([]entity.Todo, error)
	UpdateTodo(entity.Todo) error
	DeleteTodo(id uint) error
}

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoUseCase {
	return &TodoService{repo: repo}
}

func (ts *TodoService) CreateTodo(todo entity.Todo) error {
	return ts.repo.Create(todo)
}

func (ts *TodoService) GetTodo(id uint) (entity.Todo, error) {
	return ts.repo.Get(id)
}

func (ts *TodoService) GetAllTodos() ([]entity.Todo, error) {
	return ts.repo.FindAll()
}

func (ts *TodoService) GetAllTodosByUserID(userID uint) ([]entity.Todo, error) {
	return ts.repo.FindByUser(userID)
}

func (ts *TodoService) UpdateTodo(todo entity.Todo) error {
	return ts.repo.Update(todo)
}

func (ts *TodoService) DeleteTodo(id uint) error {
	return ts.repo.Delete(id)
}
