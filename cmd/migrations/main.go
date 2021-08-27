package main

import (
	"github.com/joho/godotenv"
	"go_modules_todo/internal/database"
	"go_modules_todo/internal/models"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db := database.GetDB()

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
}
