package handlers

import (
	"github.com/gin-gonic/gin"
	"go_modules_todo/internal/managers"
	"go_modules_todo/internal/requests"
	"go_modules_todo/internal/utils"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	createTodoRequest, errorsMap := requests.ValidateCreateTodoRequest(c)
	utils.CheckValidationResults(c, errorsMap)

	createdTodo := managers.CreateTodo(createTodoRequest.Description)

	c.JSON(http.StatusCreated, gin.H{"data": createdTodo})
}
