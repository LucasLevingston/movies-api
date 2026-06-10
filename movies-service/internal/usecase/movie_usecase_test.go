package usecase_test

import (
	"context"
	"errors"
	"testing"

	"movies-api/movies-service/internal/domain"
	"movies-api/movies-service/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (mockRepo *mockRepository) FindAll(ctx context.Context) ([]domain.Movie, error) {
	args := mockRepo.Called(ctx)
	return args.Get(0).([]domain.Movie), args.Error(1)
}

func (mockRepo *mockRepository) FindByID(ctx context.Context, id string) (*domain.Movie, error) {
	args := mockRepo.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Movie), args.Error(1)
}

func (mockRepo *mockRepository) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	args := mockRepo.Called(ctx, movie)
	return args.Get(0).(*domain.Movie), args.Error(1)
}

func (mockRepo *mockRepository) Delete(ctx context.Context, id string) error {
	args := mockRepo.Called(ctx, id)
	return args.Error(0)
}

func TestGetAll(t *testing.T) {
	repository := new(mockRepository)
	uc := usecase.NewMovieUseCase(repository, nil)

	expected := []domain.Movie{
		{ID: "507f1f77bcf86cd799439011", ExternalID: 8, Title: "Edison Kinetoscopic Record of a Sneeze (1894)", Year: "1894"},
		{ID: "507f1f77bcf86cd799439012", ExternalID: 10, Title: "La sortie des usines Lumière (1895)", Year: "1895"},
	}
	repository.On("FindAll", mock.Anything).Return(expected, nil)

	movies, err := uc.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, movies, 2)
	assert.Equal(t, "Edison Kinetoscopic Record of a Sneeze (1894)", movies[0].Title)
	repository.AssertExpectations(t)
}

func TestGetAll_Error(t *testing.T) {
	repository := new(mockRepository)
	uc := usecase.NewMovieUseCase(repository, nil)

	repository.On("FindAll", mock.Anything).Return([]domain.Movie{}, errors.New("db error"))

	_, err := uc.GetAll(context.Background())
	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
	repository.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	repository := new(mockRepository)
	uc := usecase.NewMovieUseCase(repository, nil)

	id := "507f1f77bcf86cd799439011"
	expected := &domain.Movie{ID: id, ExternalID: 8, Title: "Edison Kinetoscopic Record of a Sneeze (1894)", Year: "1894"}

	repository.On("FindByID", mock.Anything, id).Return(expected, nil)

	movie, err := uc.GetByID(context.Background(), id)
	assert.NoError(t, err)
	assert.Equal(t, expected.Title, movie.Title)
	repository.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	repository := new(mockRepository)
	uc := usecase.NewMovieUseCase(repository, nil)

	repository.On("FindByID", mock.Anything, "nonexistent").Return(nil, domain.ErrNotFound)

	_, err := uc.GetByID(context.Background(), "nonexistent")
	assert.ErrorIs(t, err, domain.ErrNotFound)
	repository.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	repository := new(mockRepository)
	uc := usecase.NewMovieUseCase(repository, nil)

	created := &domain.Movie{
		ID:         "507f1f77bcf86cd799439011",
		ExternalID: 999,
		Title:      "New Movie",
		Year:       "2024",
	}
	repository.On("Create", mock.Anything, mock.AnythingOfType("*domain.Movie")).Return(created, nil)

	movie, err := uc.Create(context.Background(), 999, "New Movie", "2024")
	assert.NoError(t, err)
	assert.Equal(t, "New Movie", movie.Title)
	assert.Equal(t, int32(999), movie.ExternalID)
	repository.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repository := new(mockRepository)
	uc := usecase.NewMovieUseCase(repository, nil)

	id := "507f1f77bcf86cd799439011"
	repository.On("Delete", mock.Anything, id).Return(nil)

	err := uc.Delete(context.Background(), id)
	assert.NoError(t, err)
	repository.AssertExpectations(t)
}

func TestDelete_NotFound(t *testing.T) {
	repository := new(mockRepository)
	uc := usecase.NewMovieUseCase(repository, nil)

	repository.On("Delete", mock.Anything, "bad-id").Return(domain.ErrNotFound)

	err := uc.Delete(context.Background(), "bad-id")
	assert.ErrorIs(t, err, domain.ErrNotFound)
	repository.AssertExpectations(t)
}
