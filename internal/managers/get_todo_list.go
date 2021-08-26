package managers

import (
	"go_modules/internal/database"
	"go_modules/internal/models"
)

func GetTodoList() []models.Todo {
	db := database.GetDB()

	var todos []models.Todo
	db.Find(&todos)

	return todos
}
