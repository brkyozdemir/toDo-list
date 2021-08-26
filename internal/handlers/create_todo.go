package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_modules/internal/managers"
	"go_modules/internal/requests"
	"go_modules/internal/utils"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	createTodoRequest, errorsMap := requests.ValidateCreateTodoRequest(c)
	fmt.Println(errorsMap)
	utils.CheckValidationResults(c, errorsMap)

	createdTodo := managers.CreateTodo(createTodoRequest.Description)

	c.JSON(http.StatusCreated, gin.H{"data": createdTodo})
}
