# ToDo App на Go

Простое RESTful-приложение для управления задачами, реализованное на языке Go с использованием чистой архитектуры.

## 🚀 Запуск

### 🔧 Локально

```bash
make run
```

### 🐳 Через Docker

```bash
make docker-build
make docker-run
```

Остановить контейнер:

```bash
make docker-stop
```

## 🧠 REST API

### Получить все задачи

```http
GET /todos
```

### Создать задачу

```http
POST /todos
Content-Type: application/json

{
  "title": "Сделать дз",
  "completed": false
}
```

### Удалить задачу

```http
DELETE /todos/{id}
```
