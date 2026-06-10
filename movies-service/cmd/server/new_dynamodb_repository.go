package main

import (
	"context"
	"log"

	dynamoAdapter "movies-api/movies-service/internal/adapters/dynamodb"
	"movies-api/movies-service/internal/ports"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func newDynamoDBRepository() ports.MovieRepository {
	endpoint := getEnv("DYNAMODB_ENDPOINT", "")
	tableName := getEnv("DYNAMODB_TABLE", "movies")
	region := getEnv("AWS_REGION", "us-east-1")

	options := []func(*config.LoadOptions) error{
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			getEnv("AWS_ACCESS_KEY_ID", "test"),
			getEnv("AWS_SECRET_ACCESS_KEY", "test"),
			"",
		)),
	}

	cfg, err := config.LoadDefaultConfig(context.Background(), options...)
	if err != nil {
		log.Fatalf("load aws config: %v", err)
	}

	clientOptions := []func(*dynamodb.Options){}
	if endpoint != "" {
		clientOptions = append(clientOptions, func(opts *dynamodb.Options) {
			opts.BaseEndpoint = aws.String(endpoint)
		})
	}

	client := dynamodb.NewFromConfig(cfg, clientOptions...)
	ensureDynamoTable(client, tableName)
	log.Printf("connected to DynamoDB (table: %s)", tableName)

	return dynamoAdapter.NewMovieRepository(client, tableName)
}

func ensureDynamoTable(client *dynamodb.Client, tableName string) {
	_, err := client.CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName: &tableName,
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("id"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("id"), KeyType: types.KeyTypeHash},
		},
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		log.Printf("create dynamodb table (may already exist): %v", err)
	}
}
