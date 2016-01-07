package filestore

import (
	"errors"
	"os"
)

//DeleteAllData deletes all your data.
func (store fileEventstore) DeleteAllData() error {
	if store.dataDirectory == "" {
		return errors.New("dataDirectory path cannot be blank. Did you call FileStore()?")
	}
	mutex.Lock()
	err := os.RemoveAll(store.dataDirectory)
	mutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}
