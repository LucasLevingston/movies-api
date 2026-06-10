package dynamodb

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"movies-api/movies-service/internal/domain"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (repository *movieRepository) Create(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	movie.ID = generateUUID()
	item := movieItem{
		ID:         movie.ID,
		ExternalID: movie.ExternalID,
		Title:      movie.Title,
		Year:       movie.Year,
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return nil, err
	}

	_, err = repository.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &repository.tableName,
		Item:      av,
	})
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func generateUUID() string {
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	randomBytes[6] = (randomBytes[6] & 0x0f) | 0x40
	randomBytes[8] = (randomBytes[8] & 0x3f) | 0x80
	return hex.EncodeToString(randomBytes[:4]) + "-" +
		hex.EncodeToString(randomBytes[4:6]) + "-" +
		hex.EncodeToString(randomBytes[6:8]) + "-" +
		hex.EncodeToString(randomBytes[8:10]) + "-" +
		hex.EncodeToString(randomBytes[10:])
}
