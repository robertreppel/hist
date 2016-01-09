# Hist: A Simple Eventstore in Go.

A data store for event sourced applications.

## Usage

```golang
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
```

Full (minimalist) event sourcing example:

```
cd examples/planets
go get
go build
./planets
```

## Design

Events are stored in files. Each aggregate type is a directory. Each aggregate instance is a file, with events appended
when they are saved. For example, given a data directory _"/data"_, a _"User"_ aggregate and a user with id _"12345"_, when an
"EmailChanged" event is saved it is appended to _"/data/events/User/12345.events"_

## Tests

Uses http://goconvey.co/. Run it to see BDD-style details about hist's business rules and behaviour.

## Production Use

Hist is considered alpha. Not recommended for production use.
