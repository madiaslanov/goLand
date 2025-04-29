package repository

import (
	"database/sql"
	"todoapp/internal/models"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetTodoByID(id int) (models.Todo, error) {
	var todo models.Todo
	err := r.db.QueryRow(`
		SELECT id, title, description, status, user_id, created_at, updated_at
		FROM todos WHERE id = $1`, id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt)
	return todo, err
}

func (r *TodoRepository) UpdateTodo(todo models.Todo) error {
	_, err := r.db.Exec(`
		UPDATE todos 
		SET title = $1, description = $2, status = $3, user_id = $4, updated_at = NOW()
		WHERE id = $5`,
		todo.Title, todo.Description, todo.Status, todo.UserID, todo.ID)
	return err
}

func (r *TodoRepository) CreateTodo(todo models.Todo) error {
	_, err := r.db.Exec(`
		INSERT INTO todos (title, description, status, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())`,
		todo.Title, todo.Description, todo.Status, todo.UserID)
	return err
}
