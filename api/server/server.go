package server

import (
	data "github.com/kindaqt/movies/api/data/movies"
	handlers "github.com/kindaqt/movies/api/handlers"
	"github.com/kindaqt/movies/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.GET("/healthz", handlers.HealthzHandler)

	// Movies
	h := handlers.Persister{
		// Store: data.NewStore("json"),
		Store: data.NewStore("psql"),
	}

	r.Group("/movies")
	{
		r.GET("/movies", h.GetMoviesHandler)
		r.PATCH("/movies/watched", h.UpdateWatchedHandler)
	}

	return r
}
