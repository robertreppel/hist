package leveldb

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/robertreppel/hist"
	"github.com/syndtr/goleveldb/leveldb"
)

//Save persists an event to a stream.
func (store levelDbEventstore) Save(streamID string, eventType string, eventData []byte) error {
	err := checkMandatoryParameters(streamID, eventType, eventData)
	if err != nil {
		return err
	}

	now := time.Now()
	eventRecord := hist.Event{Timestamp: now, StreamID: streamID, Type: eventType, Data: eventData}
	eventRecordJSON, err := json.Marshal(eventRecord)

	db, err := leveldb.OpenFile(store.dataDirectory, nil)
	failIf(err)
	defer db.Close()

	var maxVersionAllKey = []byte("$maxversion-$all")
	levelDbMutex.Lock()
	// Find the new maxVersion:
	maxVersion, err := db.Get(maxVersionAllKey, nil)
	if err != nil && strings.Contains(err.Error(), "not found") {
		err = db.Put([]byte(maxVersionAllKey), []byte("00000000000000000000"), nil)
		failIf(err)
		maxVersion = []byte("00000000000000000000")
	}
	failIf(err)
	currentMaxVersion, err := strconv.Atoi(string(maxVersion))
	failIf(err)
	newMaxVersion := fmt.Sprintf("%020d", currentMaxVersion+1)

	streamIDWithMaxVersion := streamID + "-" + newMaxVersion
	allStreamIDWithMaxVersion := "$all-" + newMaxVersion
	batch := new(leveldb.Batch)
	// Store all events to the $all stream:
	batch.Put([]byte(allStreamIDWithMaxVersion), []byte(eventRecordJSON))
	// In addition, store a key to look up the actual data from $all to the stream identified by the streamID:
	batch.Put([]byte(streamIDWithMaxVersion), []byte(allStreamIDWithMaxVersion))
	// Record the new maxVersion for the $all stream.
	// This is needed to support optimistic locking and for event stream subscribers wanting to catch up on events starting with a certain version.
	batch.Put([]byte(maxVersionAllKey), []byte(newMaxVersion))
	err = db.Write(batch, nil)
	failIf(err)

	levelDbMutex.Unlock()
	return nil
}

func checkMandatoryParameters(streamID string, eventType string, eventData []byte) error {
	if strings.TrimSpace(streamID) == "" {
		return errors.New("streamID cannot be blank")
	}
	if strings.TrimSpace(eventType) == "" {
		return errors.New("eventType cannot be blank")
	}
	if len(eventData) == 0 {
		return errors.New("eventData cannot be blank")
	}
	return nil
}

func failIf(err error) {
	if err != nil {
		panic(err)
	}
}
