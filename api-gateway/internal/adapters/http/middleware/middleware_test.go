package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"movies-api/api-gateway/internal/adapters/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func newTestEngine(handlers ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(handlers...)
	router.GET("/test", func(ginContext *gin.Context) {
		ginContext.Status(http.StatusOK)
	})
	router.OPTIONS("/test", func(ginContext *gin.Context) {
		ginContext.Status(http.StatusOK)
	})
	return router
}

// --- Recovery ---

func TestRecovery_PassThrough(t *testing.T) {
	router := newTestEngine(middleware.Recovery())
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRecovery_HandlerPanic(t *testing.T) {
	router := gin.New()
	router.Use(middleware.Recovery())
	router.GET("/panic", func(ginContext *gin.Context) {
		panic("unexpected error")
	})
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/panic", nil))
	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}

// --- Logging ---

func TestLogging_PassThrough(t *testing.T) {
	router := newTestEngine(middleware.Logging())
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

// --- CORS ---

func TestCORS_SetsHeaders(t *testing.T) {
	router := newTestEngine(middleware.CORS())
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Origin", "http://other-domain.com")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.NotEmpty(t, recorder.Header().Get("Access-Control-Allow-Origin"))
}

func TestCORS_PreflightOptions(t *testing.T) {
	router := newTestEngine(middleware.CORS())
	req := httptest.NewRequest(http.MethodOptions, "/test", nil)
	req.Header.Set("Origin", "http://other-domain.com")
	req.Header.Set("Access-Control-Request-Method", "POST")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusNoContent, recorder.Code)
}

// --- RequestID ---

func TestRequestID_GeneratesWhenAbsent(t *testing.T) {
	router := newTestEngine(middleware.RequestID())
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.NotEmpty(t, recorder.Header().Get(middleware.RequestIDHeader))
}

func TestRequestID_PropagatesExistingID(t *testing.T) {
	router := newTestEngine(middleware.RequestID())
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set(middleware.RequestIDHeader, "my-request-id")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, "my-request-id", recorder.Header().Get(middleware.RequestIDHeader))
}

// --- RateLimit ---

func TestRateLimit_AllowsNormalRequest(t *testing.T) {
	router := newTestEngine(middleware.RateLimit())
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestRateLimit_Returns429WhenExceeded(t *testing.T) {
	router := newTestEngine(middleware.RateLimitWithConfig(0, 1))
	// First request uses the single burst token
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.Equal(t, http.StatusOK, recorder.Code)
	// Second request: no tokens left (rate=0, no replenishment)
	recorder = httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.Equal(t, http.StatusTooManyRequests, recorder.Code)
}

// --- JWT ---

func TestJWT_PassThroughWhenNoSecret(t *testing.T) {
	router := newTestEngine(middleware.JWT(""))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestJWT_RejectsRequestWithoutToken(t *testing.T) {
	router := newTestEngine(middleware.JWT("secret"))
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/test", nil))
	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
}

func TestJWT_AcceptsValidToken(t *testing.T) {
	secret := "test-secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "user",
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	signed, _ := token.SignedString([]byte(secret))

	router := newTestEngine(middleware.JWT(secret))
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+signed)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestJWT_RejectsExpiredToken(t *testing.T) {
	secret := "test-secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "user",
		"exp": time.Now().Add(-time.Hour).Unix(),
	})
	signed, _ := token.SignedString([]byte(secret))

	router := newTestEngine(middleware.JWT(secret))
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer "+signed)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
}
