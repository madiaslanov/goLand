package main

import (
	"todoapp/internal/infrastructure/db"
	"todoapp/internal/interface/controller"
	"todoapp/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()
	taskRepo := db.NewTaskRepository(database)
	taskUsecase := usecase.NewTaskUseCase(taskRepo)

	router := gin.Default()

	// Раздача HTML
	router.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})

	// Регистрируем все API-роуты
	controller.NewTaskController(router, taskUsecase)

	router.Run(":8080")
}
