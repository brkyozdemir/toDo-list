package managers

import (
	"github.com/toDo-list/internal/database"
	"github.com/toDo-list/internal/models"
)

func CreateTodo(description string) models.Todo {
	db := database.GetDB()

	todo := &models.Todo{
		Description: description,
	}
	db.Create(todo)

	return *todo
}
