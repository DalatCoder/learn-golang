package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingGet good for passing parameters, good for testing
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"hello": "world",
		})
	}
}