package main

import (
	"go_modules/internal/database"
	"go_modules/internal/models"
)

func main() {
	db := database.GetDB()

	err := db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
}
