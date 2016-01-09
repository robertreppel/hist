package filestore

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/robertreppel/hist"
)

var mutex = &sync.Mutex{}

//Save persists an event for an aggregate.
func (store fileEventstore) Save(aggregateType string, aggregateID string, eventType string, eventData []byte) error {
	if !exists(store.dataDirectory) {
		return errors.New("No data directory")
	}
	err := checkMandatoryParameters(aggregateType, aggregateID, eventType, eventData)
	if err != nil {
		return err
	}

	//If no directory exists for storing instances of the aggregate type, create one:
	aggregatePath := store.eventsDirectory + "/" + aggregateType
	mutex.Lock()
	aggregateExists := exists(aggregatePath)
	if !aggregateExists {
		err = createDirectory(aggregatePath)
		if err != nil {
			mutex.Unlock()
			return err
		}
	}

	var file *os.File

	// If no file exists for this aggregate instance, create one:
	aggregateInstanceFile := aggregatePath + "/" + aggregateID + ".events"
	if !exists(aggregateInstanceFile) {
		file, err = store.createAggregate(aggregatePath, aggregateID)
		if err != nil {
			mutex.Unlock()
			return err
		}
		defer file.Close()
	}

	// Append the new event to the aggregate instance file:
	if file == nil {
		file, err = os.OpenFile(aggregateInstanceFile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			mutex.Unlock()
			return err
		}
		defer file.Close()
	}
	now := time.Now()
	eventRecord := hist.Event{Timestamp: now, Type: eventType, Data: eventData}
	eventRecordJSON, err := json.Marshal(eventRecord)
	if err != nil {
		mutex.Unlock()
		return err
	}
	encodedEventRecord := b64.StdEncoding.EncodeToString(eventRecordJSON)
	if _, err = file.WriteString(encodedEventRecord + "\n"); err != nil {
		mutex.Unlock()
		return err
	}
	file.Close()
	mutex.Unlock()
	return nil
}

func checkMandatoryParameters(aggregateType string, aggregateID string, eventType string, eventData []byte) error {
	if strings.TrimSpace(aggregateType) == "" {
		return errors.New("aggregateType cannot be blank")
	}
	if strings.TrimSpace(aggregateID) == "" {
		return errors.New("aggregateID cannot be blank")
	}
	if strings.TrimSpace(eventType) == "" {
		return errors.New("eventType cannot be blank")
	}
	if len(eventData) == 0 {
		return errors.New("eventData cannot be blank")
	}
	return nil
}
