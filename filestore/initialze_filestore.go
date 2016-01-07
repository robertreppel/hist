package filestore

import (
	"os"

	"github.com/robertreppel/hist"
)

type fileEventstore struct {
	dataDirectory   string
	eventsDirectory string
}

//FileStore stores events in the local file system.
func FileStore(path string) (hist.Eventstore, error) {
	var store fileEventstore
	mutex.Lock()
	dataDirectoryExists, err := exists(path)
	if err != nil {
		mutex.Unlock()
		return nil, err
	}
	store.dataDirectory = path
	store.eventsDirectory = store.dataDirectory + "/events"
	if !dataDirectoryExists {
		err = createDirectory(store.dataDirectory)
		if err != nil {
			mutex.Unlock()
			return nil, err
		}
	}

	eventDirectoryExists, err := exists(store.eventsDirectory)
	if err != nil {
		mutex.Unlock()
		return nil, err
	}
	if !eventDirectoryExists {
		err = createDirectory(store.eventsDirectory)
		if err != nil {
			mutex.Unlock()
			return nil, err
		}
	}
	mutex.Unlock()
	return store, nil
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
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
