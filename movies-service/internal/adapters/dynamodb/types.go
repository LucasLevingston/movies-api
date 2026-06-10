package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type movieRepository struct {
	client    *dynamodb.Client
	tableName string
}

type movieItem struct {
	ID         string `dynamodbav:"id"`
	ExternalID int32  `dynamodbav:"external_id"`
	Title      string `dynamodbav:"title"`
	Year       string `dynamodbav:"year"`
}
