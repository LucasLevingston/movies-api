package dynamodb

import (
	"movies-api/movies-service/internal/ports"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// NewMovieRepository creates a DynamoDB-backed MovieRepository.
func NewMovieRepository(client *dynamodb.Client, tableName string) ports.MovieRepository {
	return &movieRepository{client: client, tableName: tableName}
}
