package httphandler_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httphandler "movies-api/api-gateway/internal/adapters/http"
	"movies-api/api-gateway/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockMovieClient struct {
	mock.Mock
}

func (client *mockMovieClient) ListMovies(ctx context.Context) ([]domain.Movie, error) {
	args := client.Called(ctx)
	return args.Get(0).([]domain.Movie), args.Error(1)
}

func (client *mockMovieClient) GetMovie(ctx context.Context, id string) (*domain.Movie, error) {
	args := client.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Movie), args.Error(1)
}

func (client *mockMovieClient) CreateMovie(ctx context.Context, externalID int32, title, year string) (*domain.Movie, error) {
	args := client.Called(ctx, externalID, title, year)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Movie), args.Error(1)
}

func (client *mockMovieClient) DeleteMovie(ctx context.Context, id string) error {
	args := client.Called(ctx, id)
	return args.Error(0)
}

func newTestRouter(client *mockMovieClient) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := httphandler.NewMovieHandler(client)
	router.GET("/movies", handler.ListMovies)
	router.GET("/movies/:id", handler.GetMovie)
	router.POST("/movies", handler.CreateMovie)
	router.DELETE("/movies/:id", handler.DeleteMovie)
	return router
}

func TestListMovies_Success(t *testing.T) {
	client := new(mockMovieClient)
	router := newTestRouter(client)

	movies := []domain.Movie{
		{ID: "abc123", ExternalID: 8, Title: "Edison Kinetoscopic Record of a Sneeze (1894)", Year: "1894"},
		{ID: "def456", ExternalID: 10, Title: "La sortie des usines Lumière (1895)", Year: "1895"},
	}
	client.On("ListMovies", mock.Anything).Return(movies, nil)

	req := httptest.NewRequest(http.MethodGet, "/movies", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var result []httphandler.MovieResponse
	json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.Len(t, result, 2)
	client.AssertExpectations(t)
}

func TestListMovies_Error(t *testing.T) {
	client := new(mockMovieClient)
	router := newTestRouter(client)

	client.On("ListMovies", mock.Anything).Return([]domain.Movie{}, errors.New("service unavailable"))

	req := httptest.NewRequest(http.MethodGet, "/movies", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	client.AssertExpectations(t)
}

func TestGetMovie_Success(t *testing.T) {
	client := new(mockMovieClient)
	router := newTestRouter(client)

	movie := &domain.Movie{ID: "abc123", ExternalID: 8, Title: "Edison Kinetoscopic Record of a Sneeze (1894)", Year: "1894"}
	client.On("GetMovie", mock.Anything, "abc123").Return(movie, nil)

	req := httptest.NewRequest(http.MethodGet, "/movies/abc123", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	var result httphandler.MovieResponse
	json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.Equal(t, "abc123", result.ID)
	client.AssertExpectations(t)
}

func TestGetMovie_NotFound(t *testing.T) {
	client := new(mockMovieClient)
	router := newTestRouter(client)

	client.On("GetMovie", mock.Anything, "notexists").Return(nil, errors.New("movie not found"))

	req := httptest.NewRequest(http.MethodGet, "/movies/notexists", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	client.AssertExpectations(t)
}

func TestCreateMovie_Success(t *testing.T) {
	client := new(mockMovieClient)
	router := newTestRouter(client)

	created := &domain.Movie{ID: "abc123", ExternalID: 999, Title: "New Movie", Year: "2024"}
	client.On("CreateMovie", mock.Anything, int32(999), "New Movie", "2024").Return(created, nil)

	body := `{"external_id": 999, "title": "New Movie", "year": "2024"}`
	req := httptest.NewRequest(http.MethodPost, "/movies", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	var result httphandler.MovieResponse
	json.Unmarshal(recorder.Body.Bytes(), &result)
	assert.Equal(t, "New Movie", result.Title)
	client.AssertExpectations(t)
}

func TestCreateMovie_BadRequest(t *testing.T) {
	client := new(mockMovieClient)
	router := newTestRouter(client)

	body := `{"title": "Missing Fields"}`
	req := httptest.NewRequest(http.MethodPost, "/movies", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestDeleteMovie_Success(t *testing.T) {
	client := new(mockMovieClient)
	router := newTestRouter(client)

	client.On("DeleteMovie", mock.Anything, "abc123").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/movies/abc123", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	client.AssertExpectations(t)
}

func TestDeleteMovie_Error(t *testing.T) {
	client := new(mockMovieClient)
	router := newTestRouter(client)

	client.On("DeleteMovie", mock.Anything, "bad").Return(errors.New("not found"))

	req := httptest.NewRequest(http.MethodDelete, "/movies/bad", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	client.AssertExpectations(t)
}
