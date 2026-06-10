package dynamodb

import (
	"context"

	"movies-api/movies-service/internal/domain"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (repository *movieRepository) Delete(ctx context.Context, id string) error {
	output, err := repository.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName:    &repository.tableName,
		Key:          map[string]types.AttributeValue{"id": &types.AttributeValueMemberS{Value: id}},
		ReturnValues: types.ReturnValueAllOld,
	})
	if err != nil {
		return err
	}
	if len(output.Attributes) == 0 {
		return domain.ErrNotFound
	}
	return nil
}
