package server

import (
	data "github.com/kindaqt/movies/api/data/movies"
	handlers "github.com/kindaqt/movies/api/handlers"
	"github.com/kindaqt/movies/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware())
	// TODO: logger (ip)
	// TODO: sessions

	// Healthz
	r.GET("/healthz", handlers.HealthzHandler)

	// Movies
	h := handlers.Persister{
		// Store: data.NewStore("json"),
		Store: data.NewStore("psql"),
	}
	r.GET("/movies", h.GetMoviesHandler)
	r.PATCH("/movies/watched", h.UpdateWatchedHandler)
	r.DELETE("/movies/delete", h.DeleteMovieHandler)

	return r
}
