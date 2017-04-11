package leveldb

import (
	"errors"
	"strings"
	"sync"

	"github.com/robertreppel/hist"
)

var levelDbMutex = &sync.Mutex{}

type levelDbEventstore struct {
	dataDirectory string
}

//Store stores events in the local file system.
func Store(dataDirectory string) (hist.Eventstore, error) {
	if len(strings.TrimSpace(dataDirectory)) == 0 {
		return nil, errors.New("dataDirectory cannot be blank")
	}

	var store levelDbEventstore
	store.dataDirectory = dataDirectory
	return store, nil
}
