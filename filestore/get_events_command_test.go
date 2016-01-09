package filestore

import (
	"testing"

	"github.com/nu7hatch/gouuid"

	"github.com/robertreppel/hist"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGettingAggregateWhichDoesntExist(t *testing.T) {
	Convey("Given an event store", t, func() {
		dataStoreDirectory := "/tmp/hist-filestore-test-data"
		var store hist.Eventstore
		var err error
		store, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		Convey("when I get an aggregate which doesn't exist", func() {
			events, err := store.Get("UnknownAggregate", "12345")
			if err != nil {
				panic(err)
			}
			Convey("the returned list of events should be empty.", func() {
				So(events, ShouldBeEmpty)
			})
		})
	})
}

func TestMissingDataDirectory(t *testing.T) {
	Convey("Given an event store", t, func() {
		dataStoreDirectory := "/tmp/hist-filestore-test-missing-datadirectory"
		var store hist.Eventstore
		var err error
		store, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		deleteAllData(dataStoreDirectory)
		Convey("when I get an aggregate which doesn't exist", func() {
			_, err := store.Get("UnknownAggregate", "12345")
			Convey("an error should occur.", func() {
				So(err.Error(), ShouldEqual, "Missing data directory")
			})
		})
	})
}

func TestGettingOneEvent(t *testing.T) {
	Convey("Given an event store", t, func() {
		dataStoreDirectory := "/tmp/hist-test-filestore-data"
		var eventStore hist.Eventstore
		var err error
		eventStore, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		Convey("and an event", func() {
			expectedAggregateType := "TestGetEventAggregate"
			expectedEventType := "EventType"
			expectedEventData := "An event to get."
			id, _ := uuid.NewV4()
			aggregateID := id.String()
			data := []byte(expectedEventData)
			eventStore.Save(expectedAggregateType, aggregateID, expectedEventType, data)
			Convey("when get events is called", func() {
				Convey("then an aggregate  with the event is returned.", func() {
					events, err := eventStore.Get(expectedAggregateType, aggregateID)
					if err != nil {
						panic(err)
					}
					So(len(events), ShouldEqual, 1)
					event := events[0]
					Convey("and the type of the event is correct.", func() {
						So(event.Type, ShouldEqual, expectedEventType)
					})
					Convey("and the event data returned are correct.", func() {
						So(string(event.Data), ShouldEqual, expectedEventData)
					})
				})
			})
		})
	})
}

func TestGettingMoreThanOneEvent(t *testing.T) {
	Convey("Given an event store", t, func() {
		dataStoreDirectory := "/tmp/hist-filestore-test-data"
		var eventStore hist.Eventstore
		var err error
		eventStore, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		Convey("and three events", func() {
			expectedAggregateType := "TestGetMoreThanOneEventAggregate"
			id, _ := uuid.NewV4()
			aggregateID := id.String()
			data := []byte("{\"ShipID\":\"Seagull\",\"Location\":\"At Sea\"}")
			eventStore.Save(expectedAggregateType, aggregateID, "EventType", data)

			data = []byte("Second \n event.")
			eventStore.Save(expectedAggregateType, aggregateID, "EventType", data)

			data = []byte("Third event.")
			eventStore.Save(expectedAggregateType, aggregateID, "EventType", data)

			Convey("when get events is called", func() {
				Convey("then an aggregate  with three events is returned.", func() {
					events, err := eventStore.Get(expectedAggregateType, aggregateID)
					if err != nil {
						panic(err)
					}
					So(len(events), ShouldEqual, 3)
				})
			})
		})
	})
}
