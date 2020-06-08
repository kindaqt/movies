package handlers

import (
	"log"
	"net/http"

	model "github.com/kindaqt/movies/api/model/movies"

	"github.com/gin-gonic/gin"
)

type Persister struct {
	Store model.Store
}

////////////////////////////
// GET Handlers
//////////////////////////
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

////////////////////////////
// PATCH Handlers
//////////////////////////
type UpdateWatchedRequest struct {
	ID    string `uri:"id" form:"id" json:"id" binding:"required"`
	Value bool   `uri:"value" form:"value" json:"value" binding:"required"`
}

// UpdateWatchedHandler updates watched value
func (p *Persister) UpdateWatchedHandler(c *gin.Context) {
	log.Println("Handlers: UpdateWatchedHandler")

	var movie UpdateWatchedRequest
	if err := c.ShouldBindUri(&movie); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update Value
	if err := p.Store.UpdateWatched(movie.ID, movie.Value); err != nil {
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

////////////////////////////
// DELETE Handlers
//////////////////////////

type DeleteRequest struct {
	ID string `uri:"id" form:"id" json:"id" binding:"required"`
}

// DeleteMovieHandler deletes a movie
func (p *Persister) DeleteMovieHandler(c *gin.Context) {
	log.Println("Handlers: DeleteMovieHandler")

	// Parse Body
	var movie DeleteRequest
	if err := c.ShouldBindUri(&movie); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete movie by ID
	if err := p.Store.DeleteMovie(movie.ID); err != nil {
		log.Println(err)
		if err.Error() == "invalid id" {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	}

	c.Status(http.StatusOK)
	return
}
