package managers

import (
	"go_modules_todo/internal/database"
	"go_modules_todo/internal/models"
)

func CreateTodo(description string) models.Todo {
	db := database.GetDB()

	todo := &models.Todo{
		Description: description,
	}
	db.Create(todo)

	return *todo
}
