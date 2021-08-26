package managers

import (
	"go_modules_todo/internal/database"
	"go_modules_todo/internal/models"
)

func GetTodoList() []models.Todo {
	db := database.GetDB()

	var todos []models.Todo
	db.Find(&todos)

	return todos
}
