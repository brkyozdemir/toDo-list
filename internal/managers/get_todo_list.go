package managers

import (
	"github.com/toDo-list/internal/database"
	"github.com/toDo-list/internal/models"
)

func GetTodoList() []models.Todo {
	db := database.GetDB()

	var todos []models.Todo
	db.Find(&todos)

	return todos
}
