package dynamodb

import (
	"context"

	"movies-api/movies-service/internal/domain"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (repository *movieRepository) FindByID(ctx context.Context, id string) (*domain.Movie, error) {
	output, err := repository.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &repository.tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}
	if output.Item == nil {
		return nil, domain.ErrNotFound
	}

	var item movieItem
	if err := attributevalue.UnmarshalMap(output.Item, &item); err != nil {
		return nil, err
	}

	return &domain.Movie{
		ID:         item.ID,
		ExternalID: item.ExternalID,
		Title:      item.Title,
		Year:       item.Year,
	}, nil
}
