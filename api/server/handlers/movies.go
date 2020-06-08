package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	model "github.com/kindaqt/movies/api/model/movies"
)

type Persister struct {
	Store model.Store
}

////////////////////////////
// GET Handler
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
// PATCH Handler
//////////////////////////

type UpdateWatchedRequest struct {
	ID    string `uri:"id" form:"id" json:"id" binding:"required,uuid"`
	Value *bool  `uri:"value" form:"value" json:"value" binding:"required"`
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

	if err := p.Store.UpdateWatched(movie.ID, *movie.Value); err != nil {
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
// DELETE Handler
//////////////////////////

type DeleteRequest struct {
	ID string `uri:"id" form:"id" json:"id" binding:"required"`
}

// DeleteMovieHandler deletes a movie
func (p *Persister) DeleteMovieHandler(c *gin.Context) {
	log.Println("Handlers: DeleteMovieHandler")

	var movie DeleteRequest
	if err := c.ShouldBindUri(&movie); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Delete movie by ID
	if err := p.Store.DeleteMovie(movie.ID); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
	return
}

////////////////////////////
// CREATE Handler
//////////////////////////

type CreateRequest struct {
	Title string `uri:"title" form:"title" json:"title" binding:"required"`
}

func (p *Persister) CreateMovieHandler(c *gin.Context) {
	log.Println("Handlers: CreateMovieHandler")

	var request CreateRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create movie
	movie := model.Movie{Title: request.Title}
	if err := p.Store.CreateMovie(&movie); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"movie": movie})
	return
}
