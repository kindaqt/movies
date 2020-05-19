package server

import "github.com/gin-gonic/gin"

func Router() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})
	r.Run()
}
