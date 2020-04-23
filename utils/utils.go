package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WriteResponse is a Helper function to write response from server
func WriteResponse(err error, data interface{}, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "success",
		"result":  data,
	})
}
