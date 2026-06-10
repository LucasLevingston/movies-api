package httphandler

import (
	"net/http"

	"movies-api/api-gateway/internal/ports"

	"github.com/gin-gonic/gin"
)

// MovieHandler handles HTTP requests for movies.
type MovieHandler struct {
	client ports.MovieClient
}

// NewMovieHandler creates a new MovieHandler.
func NewMovieHandler(client ports.MovieClient) *MovieHandler {
	return &MovieHandler{client: client}
}

// MovieResponse is the movie DTO returned by the API.
type MovieResponse struct {
	ID         string `json:"id" example:"507f1f77bcf86cd799439011"`
	ExternalID int32  `json:"external_id" example:"8"`
	Title      string `json:"title" example:"Edison Kinetoscopic Record of a Sneeze (1894)"`
	Year       string `json:"year" example:"1894"`
}

// CreateMovieRequest is the payload for creating a movie.
type CreateMovieRequest struct {
	ExternalID int32  `json:"external_id" binding:"required" example:"999"`
	Title      string `json:"title" binding:"required" example:"My Movie"`
	Year       string `json:"year" binding:"required" example:"2024"`
}

// ErrorResponse wraps an error message.
type ErrorResponse struct {
	Error string `json:"error" example:"movie not found"`
}

// SuccessResponse wraps a success message.
type SuccessResponse struct {
	Message string `json:"message" example:"movie deleted successfully"`
}

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
