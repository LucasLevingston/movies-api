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
func (handler *MovieHandler) GetMovie(ginContext *gin.Context) {
	movie, err := handler.client.GetMovie(ginContext.Request.Context(), ginContext.Param("id"))
	if err != nil {
		ginContext.JSON(http.StatusNotFound, ErrorResponse{Error: err.Error()})
		return
	}
	ginContext.JSON(http.StatusOK, toMovieResponse(movie))
}
