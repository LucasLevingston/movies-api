package httphandler

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
