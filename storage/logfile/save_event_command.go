package logfile

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"

	"path/filepath"

	"github.com/robertreppel/hist"
)

//Save persists an event for an aggregate.
func (store fileEventstore) Save(streamID string, eventType string, eventData []byte) error {
	err := checkMandatoryParameters(streamID, eventType, eventData)
	if err != nil {
		return err
	}

	var file *os.File

	// Append the new event to the aggregate instance file:
	mutex.Lock()

	if exists(filepath.Join(store.dataDirectory, "eventlog.dat")) {
		file, err = os.OpenFile(filepath.Join(store.dataDirectory, "eventlog.dat"), os.O_APPEND|os.O_WRONLY, 0600)
	} else {
		file, err = os.Create(filepath.Join(store.dataDirectory, "eventlog.dat"))
	}

	if err != nil {
		mutex.Unlock()
		return err
	}
	defer file.Close()
	now := time.Now()
	eventRecord := hist.Event{Timestamp: now, StreamID: streamID, Type: eventType, Data: eventData}
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
