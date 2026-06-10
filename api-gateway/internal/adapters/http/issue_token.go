package httphandler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// IssueToken returns a handler that mints a signed JWT for demo/testing.
func IssueToken(secretKey string) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "user",
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"iat": time.Now().Unix(),
		})
		signed, err := token.SignedString([]byte(secretKey))
		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
			return
		}
		ginContext.JSON(http.StatusOK, gin.H{"token": signed})
	}
}
