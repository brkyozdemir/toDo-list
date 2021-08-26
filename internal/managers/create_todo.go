package managers

import (
	"go_modules/internal/database"
	"go_modules/internal/models"
)

func CreateTodo(description string) models.Todo {
	db := database.GetDB()

	todo := &models.Todo{
		Description: description,
	}
	db.Create(todo)

	return *todo
}
