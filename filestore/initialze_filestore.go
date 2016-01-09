package filestore

import (
	"errors"
	"os"
	"strings"

	"github.com/robertreppel/hist"
)

type fileEventstore struct {
	dataDirectory   string
	eventsDirectory string
}

//FileStore stores events in the local file system.
func FileStore(dataDirectory string) (hist.Eventstore, error) {
	if len(strings.TrimSpace(dataDirectory)) == 0 {
		return nil, errors.New("dataDirectory cannot be blank")
	}
	var store fileEventstore
	mutex.Lock()
	store.dataDirectory = dataDirectory
	store.eventsDirectory = store.dataDirectory + "/events"
	if !exists(dataDirectory) {
		err := createDirectory(store.dataDirectory)
		if err != nil {
			mutex.Unlock()
			return nil, err
		}
	}

	if !exists(store.eventsDirectory) {
		err := createDirectory(store.eventsDirectory)
		if err != nil {
			mutex.Unlock()
			return nil, err
		}
	}
	mutex.Unlock()
	return store, nil
}

// exists returns whether the given file or directory exists or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func createDirectory(path string) error {
	err := os.Mkdir(path, 0777)
	if err != nil {
		return err
	}
	return nil
}

func (store *fileEventstore) createAggregate(path string, aggregateID string) (*os.File, error) {
	file, err := os.Create(path + "/" + aggregateID + ".events")
	if err != nil {
		return nil, err
	}
	return file, nil
}
