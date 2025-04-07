package repository

import "todoapp/internal/entity"

type TaskRepository interface {
	GetAll() ([]entity.Task, error)
	Create(task entity.Task) (entity.Task, error)
	Update(id uint, task entity.Task) error
	Delete(id uint) error
}
