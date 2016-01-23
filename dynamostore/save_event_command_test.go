package dynamostore

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSaveEvent(t *testing.T) {
	eventstoreTable := "TestSaveEventTable"
	region := "us-west-2"
	endpoint := "http://localhost:8000"
	// endpoint := ""
	Convey("Given a dynamodb event store", t, func() {
		deleteTable(eventstoreTable)
		Create(eventstoreTable, region, endpoint)
		store, err := DynamoStore(eventstoreTable, region, endpoint)
		if err != nil {
			panic(err)
		}
		Convey("when an event is saved and retrieved", func() {
			now := time.Now().UnixNano()
			err = store.Save("SomeAggregate", "123456", "EventType", []byte("Event data."))
			Convey("then the operation succeeds.", func() {
				So(err, ShouldBeNil)
				Convey("and the event can be successfully retrieved", func() {
					events, err := store.Get("SomeAggregate", "123456")
					So(err, ShouldBeNil)
					So(len(events), ShouldEqual, 1)
					Convey("and has an aggregate type", func() {
						So(events[0].Type, ShouldEqual, "SomeAggregate")
					})
					Convey("and event data are correct", func() {
						So(string(events[0].Data), ShouldEqual, "Event data.")
					})
					Convey("and there is a timestamp", func() {
						So(events[0].Timestamp.UnixNano(), ShouldBeGreaterThan, now)
					})
					// Convey("and there is an event type", func() {
					// 	So(events[0].EventType, ShouldEqual, "SomeAggregate")
					// })
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
