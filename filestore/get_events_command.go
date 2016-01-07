package filestore

import (
	"bufio"
	b64 "encoding/base64"
	"encoding/json"
	"errors"

	"github.com/robertreppel/hist"
	// "log"
	"os"
)

//Get events gets the events for an aggregate
func (store fileEventstore) Get(aggregateType string, aggregateID string) ([]hist.Event, error) {
	path := store.eventsDirectory + "/" + aggregateType + "/" + aggregateID + ".events"
	mutex.Lock()
	aggregateExists, err := exists(path)
	if err != nil {
		mutex.Unlock()
		return nil, err
	}
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
	var events []hist.Event
	encodedMessage, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		// log.Printf("DEBUG hist.Get - line: '%v'\n", string(encodedMessage))
		// log.Printf("DEBUG hist.Get - encodedMessage: '%v'\n", string(encodedMessage))
		messageBytes, err := b64.StdEncoding.DecodeString(string(encodedMessage))
		if err != nil {
			mutex.Unlock()
			panic(err)
		}
		// log.Printf("DEBUG hist.Get - messageBytes: '%v'\n", string(messageBytes))
		var event hist.Event
		if len(messageBytes) == 0 {
			mutex.Unlock()
			return events, nil
		}
		err = json.Unmarshal(messageBytes, &event)
		if err != nil {
			errToReturn := errors.New("Failed unmarshalling to event for '" + aggregateType + ":" + aggregateID + "' - " + err.Error())
			mutex.Unlock()
			panic(errToReturn)
		}
		events = append(events, event)
		encodedMessage, isPrefix, err = r.ReadLine()
	}
	mutex.Unlock()
	return events, nil
}
