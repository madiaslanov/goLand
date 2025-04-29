package service

import (
	"todoapp/internal/models"
	"todoapp/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(r *repository.TodoRepository) *TodoService {
	return &TodoService{repo: r}
}

func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	return s.repo.GetAllTodos()
}

func (s *TodoService) GetTodoByID(id int) (*models.Todo, error) {
	return s.repo.GetTodoByID(id)
}

func (s *TodoService) CreateTodo(todo models.Todo) error {
	return s.repo.CreateTodo(todo)
}

func (s *TodoService) UpdateTodo(todo models.Todo) error {
	return s.repo.UpdateTodo(todo)
}

func (s *TodoService) DeleteTodo(id int) error {
	return s.repo.DeleteTodo(id)
}
