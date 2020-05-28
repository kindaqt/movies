package handlers

import (
	data "hello/server/data/movies"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMoviesHandler handles requests to get movies
func GetMoviesHandler(c *gin.Context) {
	jsonData, err := data.GetMovies()
	if err != nil {
		log.Println(err)
		c.Status(500)
		return
	}
	c.JSON(200, jsonData)
	return
}

// UpdateWatchedHandler updates watched value
func UpdateWatchedHandler(c *gin.Context) {
	// Parse Body
	var body struct {
		ID    string `json:"id"`
		Value bool   `json:"value"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Value
	if err := data.UpdateWatched(body.ID, body.Value); err != nil {
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
