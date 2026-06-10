//go:build integration

package mongodb_test

import (
	"context"
	"testing"

	ourmongo "movies-api/movies-service/internal/adapters/mongodb"
	"movies-api/movies-service/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupMongoDB(t *testing.T) (*mongodriver.Database, func()) {
	t.Helper()
	ctx := context.Background()

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo:7.0",
			ExposedPorts: []string{"27017/tcp"},
			WaitingFor:   wait.ForLog("Waiting for connections"),
		},
		Started: true,
	})
	require.NoError(t, err)

	host, err := container.Host(ctx)
	require.NoError(t, err)

	port, err := container.MappedPort(ctx, "27017")
	require.NoError(t, err)

	uri := "mongodb://" + host + ":" + port.Port()
	client, err := mongodriver.Connect(ctx, options.Client().ApplyURI(uri))
	require.NoError(t, err)

	db := client.Database("testdb")

	cleanup := func() {
		client.Disconnect(ctx)
		container.Terminate(ctx)
	}

	return db, cleanup
}

func TestFindAll_Empty_Integration(t *testing.T) {
	db, cleanup := setupMongoDB(t)
	defer cleanup()

	repo := ourmongo.NewMovieRepository(db)

	movies, err := repo.FindAll(context.Background())
	require.NoError(t, err)
	assert.Empty(t, movies)
}

func TestFindAll_WithData_Integration(t *testing.T) {
	db, cleanup := setupMongoDB(t)
	defer cleanup()

	repo := ourmongo.NewMovieRepository(db)
	ctx := context.Background()

	_, err := repo.Create(ctx, &domain.Movie{ExternalID: 8, Title: "Edison Kinetoscopic Record of a Sneeze (1894)", Year: "1894"})
	require.NoError(t, err)
	_, err = repo.Create(ctx, &domain.Movie{ExternalID: 10, Title: "La sortie des usines Lumière (1895)", Year: "1895"})
	require.NoError(t, err)

	movies, err := repo.FindAll(ctx)
	require.NoError(t, err)
	assert.Len(t, movies, 2)
}

func TestCreate_Integration(t *testing.T) {
	db, cleanup := setupMongoDB(t)
	defer cleanup()

	repo := ourmongo.NewMovieRepository(db)

	movie, err := repo.Create(context.Background(), &domain.Movie{
		ExternalID: 999,
		Title:      "New Film",
		Year:       "2024",
	})
	require.NoError(t, err)
	assert.NotEmpty(t, movie.ID)
	assert.Len(t, movie.ID, 24)
	assert.Equal(t, "New Film", movie.Title)
	assert.Equal(t, int32(999), movie.ExternalID)
}

func TestFindByID_Integration(t *testing.T) {
	db, cleanup := setupMongoDB(t)
	defer cleanup()

	repo := ourmongo.NewMovieRepository(db)
	ctx := context.Background()

	created, err := repo.Create(ctx, &domain.Movie{ExternalID: 8, Title: "Sneeze Film", Year: "1894"})
	require.NoError(t, err)

	found, err := repo.FindByID(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, found.ID)
	assert.Equal(t, "Sneeze Film", found.Title)
	assert.Equal(t, int32(8), found.ExternalID)
}

func TestFindByID_NotFound_Integration(t *testing.T) {
	db, cleanup := setupMongoDB(t)
	defer cleanup()

	repo := ourmongo.NewMovieRepository(db)

	_, err := repo.FindByID(context.Background(), "507f1f77bcf86cd799439011")
	assert.ErrorIs(t, err, domain.ErrNotFound)
}

func TestFindByID_InvalidID_Integration(t *testing.T) {
	db, cleanup := setupMongoDB(t)
	defer cleanup()

	repo := ourmongo.NewMovieRepository(db)

	_, err := repo.FindByID(context.Background(), "not-a-valid-objectid")
	assert.Error(t, err)
	assert.NotErrorIs(t, err, domain.ErrNotFound)
}

func TestDelete_Integration(t *testing.T) {
	db, cleanup := setupMongoDB(t)
	defer cleanup()

	repo := ourmongo.NewMovieRepository(db)
	ctx := context.Background()

	created, err := repo.Create(ctx, &domain.Movie{ExternalID: 1, Title: "To Delete", Year: "2024"})
	require.NoError(t, err)

	err = repo.Delete(ctx, created.ID)
	assert.NoError(t, err)

	_, err = repo.FindByID(ctx, created.ID)
	assert.ErrorIs(t, err, domain.ErrNotFound)
}

func TestDelete_NotFound_Integration(t *testing.T) {
	db, cleanup := setupMongoDB(t)
	defer cleanup()

	repo := ourmongo.NewMovieRepository(db)

	err := repo.Delete(context.Background(), "507f1f77bcf86cd799439011")
	assert.ErrorIs(t, err, domain.ErrNotFound)
}
