package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/toDo-list/internal/managers"
	"net/http"
)

func GetTodoList(c *gin.Context) {
	todos := managers.GetTodoList()

	c.JSON(http.StatusOK, todos)
}
