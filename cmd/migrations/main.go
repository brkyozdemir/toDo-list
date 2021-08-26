package main

import (
	"go_modules_todo/internal/database"
	"go_modules_todo/internal/models"
)

func main() {
	db := database.GetDB()

	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
}
