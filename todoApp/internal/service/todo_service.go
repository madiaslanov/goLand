package service

import (
	"todoApp/internal/models"
	"todoApp/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(r *repository.TodoRepository) *TodoService {
	return &TodoService{repo: r}
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) CreateTodo(todo models.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) DeleteTodo(id int) error {
	return s.repo.Delete(id)
}
