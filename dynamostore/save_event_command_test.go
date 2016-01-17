package dynamostore

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSaveEvent(t *testing.T) {
	eventstoreTable := "TestSaveEventTable"
	region := "us-west-2"
	endpoint := "http://localhost:8000"
	//endpoint := ""
	deleteTable(eventstoreTable)
	Create(eventstoreTable, region, endpoint)
	store, err := DynamoStore(eventstoreTable, region, endpoint)
	if err != nil {
		panic(err)
	}
	Convey("Given an event for an aggregate", t, func() {
		Convey("when event is saved", func() {
			err = store.Save("AggregateType", "1234", "EventType", []byte("Event data."))
			Convey("then the operation succeeds.", func() {
				So(err, ShouldBeNil)
				Convey("and the event can be successfully retrieved.", func() {
					_, err := store.Get("AggregateType", "1234")
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func deleteTable(tableName string) {
	config := aws.NewConfig().WithRegion("us-west-2")
	sess := session.New(config)
	db := dynamodb.New(sess)
	db.Endpoint = "http://localhost:8000"
	params := &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	}
	db.DeleteTable(params)

}
