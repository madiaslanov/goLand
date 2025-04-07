package db

import (
	"gorm.io/gorm"
	"todoapp/internal/entity"
)

type taskGormRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskGormRepository {
	return &taskGormRepository{db: db}
}

func (r *taskGormRepository) GetAll() ([]entity.Task, error) {
	var tasks []entity.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskGormRepository) Create(task entity.Task) (entity.Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *taskGormRepository) Update(id uint, task entity.Task) error {
	return r.db.Model(&entity.Task{}).Where("id = ?", id).Updates(task).Error
}

func (r *taskGormRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Task{}, id).Error
}
