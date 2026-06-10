package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMovie godoc
// @Summary      Get a movie by ID
// @Description  Returns a single movie by its MongoDB ObjectID
// @Tags         movies
// @Produce      json
// @Param        id   path      string  true  "Movie ObjectID"
// @Success      200  {object}  MovieResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /movies/{id} [get]
func (h *MovieHandler) GetMovie(c *gin.Context) {
	movie, err := h.client.GetMovie(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
