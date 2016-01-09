package main

import (
	"fmt"

	"github.com/robertreppel/hist/filestore"
)

const eventDataDirectory = "data"

func main() {
	eventStore, err := filestore.FileStore(eventDataDirectory)
	failIf(err)

	aggregateType := "Customer"
	aggregateID := "12345"
	eventType := "CustomerCreated"
	eventData := []byte("Bill Smith")
	err = eventStore.Save(aggregateType, aggregateID, eventType, eventData)
	failIf(err)

	eventHistory, err := eventStore.Get(aggregateType, aggregateID)
	failIf(err)

	fmt.Printf("Event: '%s' Event data: '%s'\n", eventHistory[0].Type, string(eventHistory[0].Data))
	// Output: Event: 'CustomerCreated' Event data: 'Bill Smith'
}

func failIf(err error) {
	if err != nil {
		panic(err)
	}
}
