package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"todoapp/internal/entity"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	db.AutoMigrate(&entity.Task{})
	return db
}
