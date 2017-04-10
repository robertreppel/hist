package filestore

import (
	"testing"

	"github.com/nu7hatch/gouuid"
	"github.com/robertreppel/hist"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMandatoryParameters(t *testing.T) {
	dataStoreDirectory := "/tmp/hist-test-filestore-data"
	var eventStore hist.Eventstore
	var err error
	eventStore, err = FileStore(dataStoreDirectory)
	if err != nil {
		panic(err)
	}
	Convey("When saving an event without a stream id", t, func() {
		err = eventStore.Save("", "EventType", []byte("Event data."))
		Convey("then an error occurs.", func() {
			So(err.Error(), ShouldEqual, "streamID cannot be blank")
		})
	})
	Convey("When saving an event without an event type", t, func() {
		err = eventStore.Save("SomeAggregate"+"-"+"1234", "", []byte("Event data."))
		Convey("then an error occurs.", func() {
			So(err.Error(), ShouldEqual, "eventType cannot be blank")
		})
	})
	Convey("When saving an event without any data", t, func() {
		aggregateType := "GravelAggregate"
		id, _ := uuid.NewV4()
		aggregateID := id.String()
		var data []byte
		err = eventStore.Save(aggregateType+"-"+aggregateID, "EventType", data)
		Convey("then an error occurs.", func() {
			So(err, ShouldNotBeNil)
		})
	})
}

func TestSavingToMissingDataDirectory(t *testing.T) {
	Convey("Given an event store with a data directory that doesn't exist", t, func() {
		dataStoreDirectory := "/tmp/hist-filestore-test-saving-missing-directory"
		var eventStore hist.Eventstore
		var err error
		eventStore, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		deleteAllData(dataStoreDirectory)
		Convey("saving an event should not succeed", func() {
			err = eventStore.Save("SomeAggregate"+"-"+"12356", "EventType", []byte("event data."))
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "open /tmp/hist-filestore-test-saving-missing-directory/eventlog.dat: no such file or directory")
		})
	})
}

func TestStoringNewEvent(t *testing.T) {
	Convey("Given an event store", t, func() {
		dataStoreDirectory := "/tmp/hist-test-filestore-data"
		var eventStore hist.Eventstore
		var err error
		eventStore, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		Convey("and an event which doesn't exist yet", func() {
			aggregateType := "TestSaveEventAggregate"
			id, _ := uuid.NewV4()
			aggregateID := id.String()
			data := []byte("Here's a test event.")
			Convey("when save event is called", func() {
				eventStore.Save(aggregateType+"-"+aggregateID, "EventType", data)
				Convey("then an aggregate  with the event exists.", func() {
					events, err := eventStore.Get(aggregateType + "-" + aggregateID)
					if err != nil {
						panic(err)
					}
					So(len(events) > 0, ShouldBeTrue)
				})
			})
		})
	})
}
