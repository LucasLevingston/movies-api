package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteMovie godoc
// @Summary      Delete a movie
// @Description  Removes a movie from the database by its MongoDB ObjectID
// @Tags         movies
// @Produce      json
// @Param        id   path      string  true  "Movie ObjectID"
// @Success      200  {object}  SuccessResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /movies/{id} [delete]
func (handler *MovieHandler) DeleteMovie(ginContext *gin.Context) {
	if err := handler.client.DeleteMovie(ginContext.Request.Context(), ginContext.Param("id")); err != nil {
		ginContext.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	ginContext.JSON(http.StatusOK, SuccessResponse{Message: "movie deleted successfully"})
}
