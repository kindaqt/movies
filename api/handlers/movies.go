package handlers

import (
	"log"
	"net/http"

	data "github.com/kindaqt/movies/api/data/movies"

	"github.com/gin-gonic/gin"
)

type Persister struct {
	Store data.Store
}

// GetMoviesHandler handles requests to get movies
func (p *Persister) GetMoviesHandler(c *gin.Context) {
	log.Println("Handlers: GetMoviesHandler")
	jsonData, err := p.Store.GetMovies()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return
	}
	c.JSON(200, jsonData)
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
			c.Status(500)
		}
		return
	}

	c.Status(200)
	return
}
