package service

import (
	"todoapp/internal/models"
	"todoapp/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetTodoByID(id int) (models.Todo, error) {
	return s.repo.GetTodoByID(id)
}

func (s *TodoService) UpdateTodo(todo models.Todo) error {
	return s.repo.UpdateTodo(todo)
}

func (s *TodoService) CreateTodo(todo models.Todo) error {
	return s.repo.CreateTodo(todo)
}
