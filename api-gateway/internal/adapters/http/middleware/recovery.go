package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(ginContext *gin.Context, recovered any) {
		log.Printf("[RECOVERY] panic: %v", recovered)
		ginContext.AbortWithStatus(http.StatusInternalServerError)
	})
}
