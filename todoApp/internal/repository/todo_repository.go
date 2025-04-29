package repository

import (
	"database/sql"
	"todoApp/internal/models"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (r *TodoRepository) Create(todo models.Todo) error {
	_, err := r.db.Exec("INSERT INTO todos (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
	return err
}

func (r *TodoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
