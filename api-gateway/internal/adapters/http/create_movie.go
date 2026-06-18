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
func (handler *MovieHandler) CreateMovie(ginContext *gin.Context) {
	var req CreateMovieRequest
	if err := ginContext.ShouldBindJSON(&req); err != nil {
		ginContext.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	movie, err := handler.client.CreateMovie(ginContext.Request.Context(), req.ExternalID, req.Title, req.Year)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	ginContext.JSON(http.StatusCreated, toMovieResponse(movie))
}
