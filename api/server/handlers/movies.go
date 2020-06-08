package handlers

import (
	"fmt"
	"log"
	"net/http"

	model "github.com/kindaqt/movies/api/model/movies"

	"github.com/gin-gonic/gin"
)

type Persister struct {
	Store model.Store
}

// GetMoviesHandler handles requests to get movies
func (p *Persister) GetMoviesHandler(c *gin.Context) {
	log.Println("Handlers: GetMoviesHandler")
	jsonData, err := p.Store.GetMovies()
	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, jsonData)
	return
}

// UpdateWatchedHandler updates watched value
func (p *Persister) UpdateWatchedHandler(c *gin.Context) {
	log.Println("Handlers: UpdateWatchedHandler")
	// Parse Body
	var body struct {
		ID    string `json:"id"`
		Value bool   `json:"value"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Value
	if err := p.Store.UpdateWatched(body.ID, body.Value); err != nil {
		log.Println(err)
		if err.Error() == "invalid id" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusOK)
	return
}

// DeleteMovieHandler deletes a movie
func (p *Persister) DeleteMovieHandler(c *gin.Context) {
	log.Println("Handlers: DeleteMovieHandler")

	// Parse Body
	var body struct {
		ID string `json:"id"`
	}

	fmt.Println(body)

	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete movie by ID
	if err := p.Store.DeleteMovie(body.ID); err != nil {
		if err.Error() == "invalid id" {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	}

	c.Status(http.StatusOK)
	return
}
