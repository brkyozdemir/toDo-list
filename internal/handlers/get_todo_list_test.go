package handlers

import (
	"github.com/joho/godotenv"
	"go_modules_todo/internal/models"
	"go_modules_todo/internal/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestGetTodoList(t *testing.T) {
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

	todoDummyArray := []models.Todo{
		{
			Description: "description1",
		},
		{
			Description: "description2",
		},
	}

	db.Create(todoDummyArray)

	var todos []models.Todo
	db.Find(&todos)

	if len(todos) != len(todoDummyArray) {
		t.Fatalf("Expected %v, got %v", len(todoDummyArray), len(todos))
	}

	if todoDummyArray[0].Description != "description1" {
		t.Fatalf("Expected %v, got %v", "description1", todoDummyArray[0].Description)
	}

	if todoDummyArray[1].Description != "description2" {
		t.Fatalf("Expected %v, got %v", "description2", todoDummyArray[1].Description)
	}

	utils.RefreshDatabase()
}
