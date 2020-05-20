package server

import "github.com/gin-gonic/gin"

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/healthz", healthzHandler)
	return r
}

func healthzHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "healthy",
	})
	// c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
