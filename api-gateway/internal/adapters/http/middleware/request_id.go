package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

const RequestIDHeader = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		requestID := ginContext.GetHeader(RequestIDHeader)
		if requestID == "" {
			raw := make([]byte, 16)
			rand.Read(raw)
			requestID = hex.EncodeToString(raw)
		}
		ginContext.Header(RequestIDHeader, requestID)
		ginContext.Set(RequestIDHeader, requestID)
		ginContext.Next()
	}
}
