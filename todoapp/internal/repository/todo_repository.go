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

func (r *TodoRepository) GetTodoByID(id int) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.QueryRow(`
		SELECT id, title, description, status, user_id, created_at, updated_at
		FROM todos WHERE id = $1`, id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &todo, nil
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

func (r *TodoRepository) GetAllTodos() ([]models.Todo, error) {
	rows, err := r.db.Query(`
		SELECT id, title, description, status, user_id, created_at, updated_at
		FROM todos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) DeleteTodo(id int) error {
	_, err := r.db.Exec(`DELETE FROM todos WHERE id = $1`, id)
	return err
}

func (r *TodoRepository) CreateHistoryRecord(todoID int, status string) error {
	_, err := r.db.Exec(`
		INSERT INTO todo_history (todo_id, status)
		VALUES ($1, $2)`,
		todoID, status)
	return err
}

func (r *TodoRepository) GetTodoHistory(todoID int) ([]models.TodoHistory, error) {
	rows, err := r.db.Query(`
		SELECT id, todo_id, status, updated_at
		FROM todo_history
		WHERE todo_id = $1
		ORDER BY updated_at DESC`, todoID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.TodoHistory
	for rows.Next() {
		var record models.TodoHistory
		if err := rows.Scan(&record.ID, &record.TodoID, &record.Status, &record.UpdatedAt); err != nil {
			return nil, err
		}
		history = append(history, record)
	}

	return history, nil
}
func (r *TodoRepository) AddTodoHistory(todoID int, status string) error {
	_, err := r.db.Exec(`
        INSERT INTO todo_history (todo_id, status, updated_at)
        VALUES ($1, $2, NOW())`, todoID, status)
	return err
}
