package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todoApp/internal/models"
	"todoApp/internal/service"

	"github.com/gorilla/mux"
)

type TodoController struct {
	service *service.TodoService
}

func NewTodoController(s *service.TodoService) *TodoController {
	return &TodoController{service: s}
}

func (c *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := c.service.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err = c.service.CreateTodo(todo)
	if err != nil {
		http.Error(w, "Failed to create todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = c.service.DeleteTodo(id)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
