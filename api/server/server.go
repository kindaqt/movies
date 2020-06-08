package server

import (
	model "github.com/kindaqt/movies/api/model/movies"
	handlers "github.com/kindaqt/movies/api/server/handlers"
	"github.com/kindaqt/movies/api/server/middleware"

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
		Store: model.NewStore("psql"),
	}
	r.GET("/movies", h.GetMoviesHandler)
	r.PATCH("/movies/:id/watched/:value", h.UpdateWatchedHandler)
	r.DELETE("/movies/:id", h.DeleteMovieHandler)
	// r.POST("/movies/:id", h.CreateMovieHandler)

	return r
}
