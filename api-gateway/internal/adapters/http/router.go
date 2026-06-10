package httphandler

import (
	"net/http"

	"movies-api/api-gateway/internal/adapters/http/middleware"
	"movies-api/api-gateway/internal/ports"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// NewRouter creates a Gin router with all middleware and movie endpoints wired up.
func NewRouter(client ports.MovieClient, jwtSecret string) *gin.Engine {
	router := gin.New()

	router.Use(middleware.Recovery())
	router.Use(middleware.Logging())
	router.Use(middleware.CORS())
	router.Use(middleware.RequestID())
	router.Use(middleware.RateLimit())
	router.Use(middleware.JWT(jwtSecret))

	handler := NewMovieHandler(client)

	router.GET("/movies", handler.ListMovies)
	router.GET("/movies/:id", handler.GetMovie)
	router.POST("/movies", handler.CreateMovie)
	router.DELETE("/movies/:id", handler.DeleteMovie)

	router.GET("/health", func(ginContext *gin.Context) {
		ginContext.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	if jwtSecret != "" {
		router.POST("/token", IssueToken(jwtSecret))
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
