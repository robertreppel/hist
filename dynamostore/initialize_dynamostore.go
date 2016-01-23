package dynamostore

import (
	"errors"
	"strings"

	"github.com/robertreppel/hist"
)

type dynamoEventstore struct {
	tableName string
	region    string
	endpoint  string
}

//DynamoStore stores events in AWS DynamoDB.
func DynamoStore(tableName string, region string, endpoint string) (hist.Eventstore, error) {
	if len(strings.TrimSpace(tableName)) == 0 {
		return nil, errors.New("tableName cannot be blank")
	}

	if len(strings.TrimSpace(region)) == 0 {
		return nil, errors.New("region cannot be blank")
	}

	var store dynamoEventstore
	store.region = region
	store.tableName = tableName
	store.endpoint = endpoint
	return store, nil
}
