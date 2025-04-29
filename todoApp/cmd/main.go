package main

import (
	"log"
	"net/http"
	"todoApp/internal/config"
	"todoApp/internal/controller"
	"todoApp/internal/middleware"
	"todoApp/internal/repository"
	"todoApp/internal/service"

	"github.com/gorilla/mux"
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
