package filestore

import (
	"testing"

	"github.com/robertreppel/hist"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateMissingDataDirectory(t *testing.T) {
	dataStoreDirectory := "/tmp/hist-test-create-missing-directory"
	Convey("Given that no data directory exists", t, func() {
		var eventStore hist.Eventstore
		var err error
		eventStore, err = FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		eventStore.DeleteAllData()
		Convey("when the FileStore is initialized", func() {
			_, err := FileStore(dataStoreDirectory)
			if err != nil {
				panic(err)
			}
			Convey("then the data directory exists.", func() {
				dbDirectoryExists, _ := exists(dataStoreDirectory)
				So(dbDirectoryExists, ShouldBeTrue)
			})
			Convey("and the events directory exists.", func() {
				dbDirectoryExists, err := exists(dataStoreDirectory + "/events")
				if err != nil {
					panic(err)
				}
				So(dbDirectoryExists, ShouldBeTrue)
			})
		})
	})
}

var eventStore fileEventstore

func TestDeleteAllData(t *testing.T) {
	dataStoreDirectory := "/tmp/hist-test-delete-all-data-directory"
	Convey("Given that a data directory exists", t, func() {
		eventStore, err := FileStore(dataStoreDirectory)
		if err != nil {
			panic(err)
		}
		Convey("when DeleteAllData is called", func() {
			err := eventStore.DeleteAllData()
			So(err, ShouldBeNil)
			Convey("then the data directory should not exist.", func() {
				dbDirectoryExists, _ := exists(dataStoreDirectory)
				So(dbDirectoryExists, ShouldBeFalse)
			})
		})
	})
}
