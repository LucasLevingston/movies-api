package usecase_test

import (
	"context"
	"errors"
	"testing"

	"movies-api/movies-service/internal/domain"
	"movies-api/movies-service/internal/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) FindAll(ctx context.Context) ([]domain.Movie, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Movie), args.Error(1)
}

func (m *mockRepository) FindByID(ctx context.Context, id string) (*domain.Movie, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Movie), args.Error(1)
}

func (m *mockRepository) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	args := m.Called(ctx, movie)
	return args.Get(0).(*domain.Movie), args.Error(1)
}

func (m *mockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestGetAll(t *testing.T) {
	repo := new(mockRepository)
	uc := usecase.NewMovieUseCase(repo)

	expected := []domain.Movie{
		{ID: primitive.NewObjectID(), ExternalID: 8, Title: "Edison Kinetoscopic Record of a Sneeze (1894)", Year: "1894"},
		{ID: primitive.NewObjectID(), ExternalID: 10, Title: "La sortie des usines Lumière (1895)", Year: "1895"},
	}
	repo.On("FindAll", mock.Anything).Return(expected, nil)

	movies, err := uc.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, movies, 2)
	assert.Equal(t, "Edison Kinetoscopic Record of a Sneeze (1894)", movies[0].Title)
	repo.AssertExpectations(t)
}

func TestGetAll_Error(t *testing.T) {
	repo := new(mockRepository)
	uc := usecase.NewMovieUseCase(repo)

	repo.On("FindAll", mock.Anything).Return([]domain.Movie{}, errors.New("db error"))

	_, err := uc.GetAll(context.Background())
	assert.Error(t, err)
	assert.EqualError(t, err, "db error")
	repo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	repo := new(mockRepository)
	uc := usecase.NewMovieUseCase(repo)

	id := primitive.NewObjectID()
	expected := &domain.Movie{ID: id, ExternalID: 8, Title: "Edison Kinetoscopic Record of a Sneeze (1894)", Year: "1894"}

	repo.On("FindByID", mock.Anything, id.Hex()).Return(expected, nil)

	movie, err := uc.GetByID(context.Background(), id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, expected.Title, movie.Title)
	repo.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	repo := new(mockRepository)
	uc := usecase.NewMovieUseCase(repo)

	repo.On("FindByID", mock.Anything, "nonexistent").Return(nil, errors.New("not found"))

	_, err := uc.GetByID(context.Background(), "nonexistent")
	assert.Error(t, err)
	repo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	repo := new(mockRepository)
	uc := usecase.NewMovieUseCase(repo)

	created := &domain.Movie{
		ID:         primitive.NewObjectID(),
		ExternalID: 999,
		Title:      "New Movie",
		Year:       "2024",
	}
	repo.On("Create", mock.Anything, mock.AnythingOfType("*domain.Movie")).Return(created, nil)

	movie, err := uc.Create(context.Background(), 999, "New Movie", "2024")
	assert.NoError(t, err)
	assert.Equal(t, "New Movie", movie.Title)
	assert.Equal(t, int32(999), movie.ExternalID)
	repo.AssertExpectations(t)
}

func TestDelete(t *testing.T) {
	repo := new(mockRepository)
	uc := usecase.NewMovieUseCase(repo)

	id := primitive.NewObjectID().Hex()
	repo.On("Delete", mock.Anything, id).Return(nil)

	err := uc.Delete(context.Background(), id)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDelete_Error(t *testing.T) {
	repo := new(mockRepository)
	uc := usecase.NewMovieUseCase(repo)

	repo.On("Delete", mock.Anything, "bad-id").Return(errors.New("not found"))

	err := uc.Delete(context.Background(), "bad-id")
	assert.Error(t, err)
	repo.AssertExpectations(t)
}
