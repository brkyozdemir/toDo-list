package handlers

import (
	"github.com/joho/godotenv"
	"go_modules/internal/models"
	"go_modules/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestCreateTodo(t *testing.T) {
	utils.CreateTestDB()
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DB_CONNECTION") + os.Getenv("TEST_DB_NAME") + os.Getenv("DB_OPTIONS")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	description := "description"
	todoTestModel := &models.Todo{
		Description: description,
	}
	db.Create(todoTestModel)

	var todo models.Todo
	db.Find(&todo)

	if todo.Description != description {
		t.Fatalf("Expected 'description', got %v", todo.Description)
	}

	utils.RefreshDatabase()
}
