package logfile

import (
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/robertreppel/hist"
)

var mutex = &sync.Mutex{}

type fileEventstore struct {
	dataDirectory string
}

//FileStore stores events in the local file system.
func FileStore(dataDirectory string) (hist.Eventstore, error) {
	if len(strings.TrimSpace(dataDirectory)) == 0 {
		return nil, errors.New("dataDirectory cannot be blank")
	}
	if !exists(dataDirectory) {
		mutex.Lock()
		err := createDirectory(dataDirectory)
		if err != nil {
			mutex.Unlock()
			return nil, err
		}
		mutex.Unlock()
	}

	var store fileEventstore
	store.dataDirectory = dataDirectory
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
