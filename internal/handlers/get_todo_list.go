package handlers

import (
	"github.com/gin-gonic/gin"
	"go_modules_todo/internal/managers"
	"net/http"
)

func GetTodoList(c *gin.Context) {
	todos := managers.GetTodoList()

	c.JSON(http.StatusOK, todos)
}
