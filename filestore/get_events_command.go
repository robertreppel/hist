package filestore

import (
	"bufio"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/robertreppel/hist"
)

//Get events gets the events for an aggregate
func (store fileEventstore) Get(streamID string) ([]hist.Event, error) {
	var events []hist.Event
	path := filepath.Join(store.dataDirectory, "eventlog.dat")

	mutex.Lock()
	aggregateExists := exists(path)
	if !aggregateExists {
		var emptyResult []hist.Event
		mutex.Unlock()
		return emptyResult, nil
	}

	inFile, err := os.Open(path)
	if err != nil {
		mutex.Unlock()
		return nil, err
	}
	defer inFile.Close()

	r := bufio.NewReader(inFile)
	encodedMessage, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		messageBytes, err := b64.StdEncoding.DecodeString(string(encodedMessage))
		if err != nil {
			mutex.Unlock()
			panic(err)
		}
		var event hist.Event
		if len(messageBytes) == 0 {
			mutex.Unlock()
			return events, nil
		}
		err = json.Unmarshal(messageBytes, &event)
		if err != nil {
			errToReturn := errors.New("Failed unmarshalling to event for '" + streamID + "' - " + err.Error())
			mutex.Unlock()
			panic(errToReturn)
		}
		if event.StreamID == streamID {
			events = append(events, event)
		}
		encodedMessage, isPrefix, err = r.ReadLine()
	}
	mutex.Unlock()
	return events, nil
}
