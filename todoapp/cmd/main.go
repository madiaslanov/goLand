package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todoapp/internal/config"
	"todoapp/internal/controller"
	"todoapp/internal/middleware"
	"todoapp/internal/repository"
	"todoapp/internal/service"
)

func main() {
	db := config.InitDB()
	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoController := controller.NewTodoController(todoService)

	r := mux.NewRouter()
	r.Use(middleware.LoggerMiddleware)

	r.HandleFunc("/todos", todoController.GetAllTodos).Methods("GET")
	r.HandleFunc("/todos", todoController.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoController.DeleteTodo).Methods("DELETE")

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
