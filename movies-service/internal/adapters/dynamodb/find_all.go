package dynamodb

import (
	"context"

	"movies-api/movies-service/internal/domain"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (repository *movieRepository) FindAll(ctx context.Context) ([]domain.Movie, error) {
	output, err := repository.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: &repository.tableName,
	})
	if err != nil {
		return nil, err
	}

	var items []movieItem
	if err := attributevalue.UnmarshalListOfMaps(output.Items, &items); err != nil {
		return nil, err
	}

	movies := make([]domain.Movie, len(items))
	for index, item := range items {
		movies[index] = domain.Movie{
			ID:         item.ID,
			ExternalID: item.ExternalID,
			Title:      item.Title,
			Year:       item.Year,
		}
	}
	return movies, nil
}
