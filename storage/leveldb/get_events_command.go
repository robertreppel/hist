package leveldb

import (
	"encoding/json"

	"github.com/robertreppel/hist"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

//Get events gets the events for a stream.
func (store levelDbEventstore) Get(streamID string) ([]hist.Event, error) {
	db, err := leveldb.OpenFile(store.dataDirectory, nil)
	failIf(err)
	defer db.Close()

	var events []hist.Event
	iter := db.NewIterator(util.BytesPrefix([]byte(streamID)), nil)
	for iter.Next() {
		dataKey := iter.Value()
		eventData, err := db.Get([]byte(dataKey), nil)
		if err != nil {
			return events, nil
		}
		var event hist.Event
		err = json.Unmarshal(eventData, &event)
		failIf(err)
		events = append(events, event)
	}
	iter.Release()
	err = iter.Error()

	return events, nil
}
