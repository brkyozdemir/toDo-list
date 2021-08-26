package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toDo-list/internal/managers"
	"github.com/toDo-list/internal/requests"
	"github.com/toDo-list/internal/utils"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	createTodoRequest, errorsMap := requests.ValidateCreateTodoRequest(c)
	fmt.Println(errorsMap)
	utils.CheckValidationResults(c, errorsMap)

	createdTodo := managers.CreateTodo(createTodoRequest.Description)

	c.JSON(http.StatusCreated, gin.H{"data": createdTodo})
}
