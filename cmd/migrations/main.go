package main

import (
	"github.com/toDo-list/internal/database"
	"github.com/toDo-list/internal/models"
)

func main() {
	db := database.GetDB()

	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
}
