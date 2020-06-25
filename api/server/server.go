package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kindaqt/movies/api/models"
	"github.com/kindaqt/movies/api/server/handlers"
	"github.com/kindaqt/movies/api/server/middleware"
)

type Server struct {
	DataStore *models.Store
}

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
		Store: models.NewStore("psql"),
	}
	fmt.Println(h)
	r.GET("/movies", h.GetMoviesHandler)
	r.POST("/movies", h.CreateMovieHandler)
	r.PATCH("/movies/:id/watched/:value", h.UpdateWatchedHandler)
	r.DELETE("/movies/:id", h.DeleteMovieHandler)

	return r
}
