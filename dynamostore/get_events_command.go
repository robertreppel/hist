package dynamostore

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/robertreppel/hist"
)

//Get events gets the events for an aggregate
func (store dynamoEventstore) Get(aggregateType string, aggregateID string) ([]hist.Event, error) {
	db := getDb(store.region, store.endpoint)
	params := &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("AggregateTypeAndId=:aggregatekey"),

		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":aggregatekey": {S: aws.String(aggregateType + ":" + aggregateID)},
		},
		TableName: aws.String(store.tableName),
	}
	resp, err := db.Query(params)
	if err != nil {
		return nil, err
	}
	log.Println(resp)

	return nil, errors.New("Not implemented.")
}
