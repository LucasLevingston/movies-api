package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		start := time.Now()
		method := ginContext.Request.Method
		path := ginContext.Request.URL.Path

		ginContext.Next()

		log.Printf("[HTTP] %s %s %d %s",
			method,
			path,
			ginContext.Writer.Status(),
			time.Since(start),
		)
	}
}
