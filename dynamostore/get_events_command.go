package dynamostore

import (
	"strconv"
	"strings"
	"time"

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
	events := []hist.Event{}
	cnt := int64(0)
	for cnt < *resp.Count {
		event := resp.Items[cnt]
		timeNano, err := strconv.Atoi(*event["Timestamp"].N)
		if err != nil {
			return nil, err
		}
		aggregateType := strings.Split(*event["AggregateTypeAndId"].S, ":")[0]
		timestamp := time.Unix(0, int64(timeNano))
		data := event["Data"].B
		newEvent := hist.Event{
			Type:      aggregateType,
			Timestamp: timestamp,
			Data:      data,
		}
		events = append(events, newEvent)
		cnt++
	}
	return events, nil
}
