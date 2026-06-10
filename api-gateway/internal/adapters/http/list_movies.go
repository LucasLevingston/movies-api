package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListMovies godoc
// @Summary      List all movies
// @Description  Returns all movies stored in the database
// @Tags         movies
// @Produce      json
// @Success      200  {array}   MovieResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /movies [get]
func (h *MovieHandler) ListMovies(c *gin.Context) {
	movies, err := h.client.ListMovies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}
