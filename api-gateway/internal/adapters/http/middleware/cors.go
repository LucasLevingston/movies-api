package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowMethods:    []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "X-Request-ID"},
		ExposeHeaders:   []string{"X-Request-ID"},
	})
}
