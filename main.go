package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go_modules/internal/handlers"
)

func main() {
	r := gin.Default()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r.POST("/todo", handlers.CreateTodo)
	r.GET("/todos", handlers.GetTodoList)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
