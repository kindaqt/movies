package server

import (
	handlers "github.com/kindaqt/movies/api/handlers"
	"github.com/kindaqt/movies/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/healthz", handlers.HealthzHandler)
	r.GET("/movies", handlers.GetMoviesHandler)
	r.PATCH("/movies/watched", handlers.GetMoviesHandler)

	return r
}
