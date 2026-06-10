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
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	if err := h.client.DeleteMovie(c.Request.Context(), c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Message: "movie deleted successfully"})
}
