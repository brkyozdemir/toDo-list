package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/toDo-list/internal/utils"
)

type CreateTodoRequest struct {
	Description string `bson:"description" json:"description"`
}

func NewCreateTodoRequest() CreateTodoRequest {
	return CreateTodoRequest{}
}

func ValidateCreateTodoRequest(c *gin.Context) (CreateTodoRequest, string) {
	requestModel := NewCreateTodoRequest()
	var errorsMap string

	// Required parameters
	err := c.BindJSON(&requestModel)
	if err != nil {
		panic(err)
	}
	if requestModel.Description == "" {
		utils.MapValidationError(&errorsMap, "description")
	}

	return requestModel, errorsMap
}
