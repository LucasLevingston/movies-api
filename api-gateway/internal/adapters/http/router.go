package httphandler

import (
	"movies-api/api-gateway/internal/ports"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(client ports.MovieClient) *gin.Engine {
	r := gin.Default()

	h := NewMovieHandler(client)

	r.GET("/movies", h.ListMovies)
	r.GET("/movies/:id", h.GetMovie)
	r.POST("/movies", h.CreateMovie)
	r.DELETE("/movies/:id", h.DeleteMovie)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
