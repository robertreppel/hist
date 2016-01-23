package dynamostore

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Save persists an event for an aggregate.
func (store dynamoEventstore) Save(aggregateType string, aggregateID string, eventType string, eventData []byte) error {
	err := checkMandatoryParameters(aggregateType, aggregateID, eventType, eventData)
	if err != nil {
		return err
	}
	db := getDb(store.region, store.endpoint)

	params := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{ // Required
			"AggregateTypeAndId": {
				S: aws.String(aggregateType + ":" + aggregateID),
			},
			"Timestamp": {
				N: aws.String(strconv.FormatInt(time.Now().UnixNano(), 10)),
			},
			"Data": {
				B: eventData,
			},
		},
		TableName: aws.String(store.tableName),
	}

	_, err = db.PutItem(params)

	if err != nil {
		return err
	}
	return nil
}

func checkMandatoryParameters(aggregateType string, aggregateID string, eventType string, eventData []byte) error {
	if strings.TrimSpace(aggregateType) == "" {
		return errors.New("aggregateType cannot be blank")
	}
	if strings.TrimSpace(aggregateID) == "" {
		return errors.New("aggregateID cannot be blank")
	}
	if strings.TrimSpace(eventType) == "" {
		return errors.New("eventType cannot be blank")
	}
	if len(eventData) == 0 {
		return errors.New("eventData cannot be blank")
	}
	return nil
}
