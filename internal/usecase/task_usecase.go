package usecase

import (
	"todoapp/internal/entity"
	"todoapp/internal/interface/repository"
)

type TaskUseCase interface {
	GetAll() ([]entity.Task, error)
	Create(task entity.Task) (entity.Task, error)
	Update(id uint, task entity.Task) error
	Delete(id uint) error
}

type taskUseCase struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(r repository.TaskRepository) TaskUseCase {
	return &taskUseCase{repo: r}
}

func (uc *taskUseCase) GetAll() ([]entity.Task, error) {
	return uc.repo.GetAll()
}

func (uc *taskUseCase) Create(task entity.Task) (entity.Task, error) {
	return uc.repo.Create(task)
}

func (uc *taskUseCase) Update(id uint, task entity.Task) error {
	return uc.repo.Update(id, task)
}

func (uc *taskUseCase) Delete(id uint) error {
	return uc.repo.Delete(id)
}
