package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	// Загружаем переменные окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Получаем данные из переменных окружения
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Формируем строку подключения
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname,
	)

	// Открываем подключение к базе данных
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}

	// Проверяем подключение
	err = db.Ping()
	if err != nil {
		log.Fatal("База не отвечает:", err)
	}

	// Сообщаем, что подключение прошло успешно
	fmt.Println("Успешное подключение к базе Supabase!")

	return db
}
