package dynamostore

import (
	"testing"

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
			err := Create(eventstoreTable, "us-west-2", "http://localhost:8000")
			Convey("then the table is created successfully.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
