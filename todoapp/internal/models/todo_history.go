package models

import "time"

type TodoHistory struct {
	ID        int       `json:"id"`
	TodoID    int       `json:"todo_id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}
