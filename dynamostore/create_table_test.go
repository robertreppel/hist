package dynamostore

import (
	"testing"

	// "github.com/nu7hatch/gouuid"

	// "github.com/robertreppel/hist"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateTable(t *testing.T) {
	eventstoreTable := "TableName"
	config := aws.NewConfig().WithRegion("us-west-2")
	sess := session.New(config)
	db := dynamodb.New(sess)
	db.Endpoint = "http://localhost:8000"
	Convey("Given that no event store DynamoDb table exists", t, func() {
		params := &dynamodb.DeleteTableInput{
			TableName: aws.String(eventstoreTable),
		}
		db.DeleteTable(params)
		Convey("when the table is created", func() {
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
				TableName: aws.String(eventstoreTable),
			}
			resp, err := db.CreateTable(params)
			if err != nil {
				panic(err)
			}
			t.Log(resp)

			Convey("then the table is created successfully.", func() {
				So(resp.String(), ShouldContainSubstring, "TableStatus: \"ACTIVE\"")
				So(resp.String(), ShouldContainSubstring, "TableName: \""+eventstoreTable+"\"")
			})
		})
	})
}
