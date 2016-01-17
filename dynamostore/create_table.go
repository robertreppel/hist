package dynamostore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Create a DynameDB event store.
func Create(tableName string, region string, endpoint string) error {
	config := aws.NewConfig().WithRegion(region)
	sess := session.New(config)
	db := dynamodb.New(sess)
	if endpoint != "" {
		db.Endpoint = endpoint
	}

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("AggregateTypeAndId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Timestamp"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("AggregateTypeAndId"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Timestamp"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(tableName),
	}
	_, err := db.CreateTable(params)
	if err != nil {
		return err
	}
	return nil
}
