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

	router.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})

	controller.NewTaskController(router, taskUsecase)

	router.Run(":8080")
}
