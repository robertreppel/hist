package filestore

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateMissingDataDirectory(t *testing.T) {
	dataStoreDirectory := "/tmp/hist-test-create-missing-directory"
	Convey("Given that no data directory exists", t, func() {
		deleteAllData(dataStoreDirectory)
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

func deleteAllData(dataDirectory string) error {
	mutex.Lock()
	err := os.RemoveAll(dataDirectory)
	mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}
