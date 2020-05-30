package handlers

import (
	"github.com/gin-gonic/gin"
)

func HealthzHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
	})
}
