package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const (
	rateLimitRequestsPerSecond rate.Limit = 10
	rateLimitBurst                        = 20
)

func RateLimit() gin.HandlerFunc {
	return RateLimitWithConfig(rateLimitRequestsPerSecond, rateLimitBurst)
}

func RateLimitWithConfig(requestsPerSecond rate.Limit, burst int) gin.HandlerFunc {
	var ipLimiters sync.Map

	return func(ginContext *gin.Context) {
		limiter, _ := ipLimiters.LoadOrStore(
			ginContext.ClientIP(),
			rate.NewLimiter(requestsPerSecond, burst),
		)
		if !limiter.(*rate.Limiter).Allow() {
			ginContext.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			return
		}
		ginContext.Next()
	}
}
