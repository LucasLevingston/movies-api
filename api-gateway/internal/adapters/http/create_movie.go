package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateMovie godoc
// @Summary      Create a movie
// @Description  Adds a new movie to the database
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        movie  body      CreateMovieRequest  true  "Movie payload"
// @Success      201    {object}  MovieResponse
// @Failure      400    {object}  ErrorResponse
// @Failure      500    {object}  ErrorResponse
// @Router       /movies [post]
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var req CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	movie, err := h.client.CreateMovie(c.Request.Context(), req.ExternalID, req.Title, req.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, movie)
}
