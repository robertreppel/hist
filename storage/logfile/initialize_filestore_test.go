package logfile

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMandatoryFilestoreDataDirectory(t *testing.T) {
	Convey("When creating a FileStore with a blank dataStoreDirectory", t, func() {
		_, err := FileStore("")
		Convey("an error should occur.", func() {
			So(err.Error(), ShouldEqual, "dataDirectory cannot be blank")

		})
	})
}

func TestCreateMissingDataDirectory(t *testing.T) {
	dataStoreDirectory := "/tmp/hist-test-filestore-create-missing-directory"
	Convey("Given that no data directory exists", t, func() {
		deleteAllData(dataStoreDirectory)
		Convey("when the FileStore is initialized", func() {
			_, err := FileStore(dataStoreDirectory)
			if err != nil {
				panic(err)
			}
			Convey("then the data directory exists.", func() {
				dbDirectoryExists := exists(dataStoreDirectory)
				So(dbDirectoryExists, ShouldBeTrue)
			})
		})
	})
}

func deleteAllData(dataDirectory string) error {
	// mutex.Lock()
	err := os.RemoveAll(dataDirectory)
	// mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}
