package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckValidationResults(c *gin.Context, err string) {
	if err != "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "The given data was invalid.",
			"error":   err,
		})
		panic(err)
	}
}
