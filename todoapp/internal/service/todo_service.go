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

	oldTodo, err := s.repo.GetTodoByID(todo.ID)
	if err != nil {
		return err
	}

	err = s.repo.UpdateTodo(todo)
	if err != nil {
		return err
	}

	if oldTodo.Status != todo.Status {
		return s.repo.AddTodoHistory(todo.ID, todo.Status)
	}

	return nil
}

func (s *TodoService) DeleteTodo(id int) error {
	return s.repo.DeleteTodo(id)
}

func (s *TodoService) UpdateTodoStatus(todo models.Todo) error {
	oldStatus := todo.Status

	err := s.repo.UpdateTodo(todo)
	if err != nil {
		return err
	}

	if oldStatus != todo.Status {
		err := s.repo.CreateHistoryRecord(todo.ID, todo.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TodoService) GetTodoHistory(todoID int) ([]models.TodoHistory, error) {
	return s.repo.GetTodoHistory(todoID)
}
