package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingHandler returns a simple pong and can be used for liveness/readiness.
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
