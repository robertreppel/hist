package dynamostore

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateTable(t *testing.T) {
	eventstoreTable := "TestCreateTableTable"
	Convey("Given that no event store DynamoDb table exists", t, func() {
		deleteTable(eventstoreTable)
		Convey("when the table is created", func() {
			err := Create(eventstoreTable, "us-west-2", "http://localhost:8000")
			Convey("then the table is created successfully.", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
