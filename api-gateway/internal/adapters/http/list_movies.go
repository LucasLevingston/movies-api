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
func (handler *MovieHandler) ListMovies(ginContext *gin.Context) {
	movies, err := handler.client.ListMovies(ginContext.Request.Context())
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	ginContext.JSON(http.StatusOK, movies)
}
