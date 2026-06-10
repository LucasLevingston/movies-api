package httphandler

import (
	"net/http"

	"movies-api/api-gateway/internal/ports"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter creates a Gin router with all movie endpoints and Swagger UI wired up.
func NewRouter(client ports.MovieClient) *gin.Engine {
	router := gin.Default()

	h := NewMovieHandler(client)

	router.GET("/movies", h.ListMovies)
	router.GET("/movies/:id", h.GetMovie)
	router.POST("/movies", h.CreateMovie)
	router.DELETE("/movies/:id", h.DeleteMovie)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
