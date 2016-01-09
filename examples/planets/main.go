package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/robertreppel/hist/filestore"
)

const eventDataDirectory = "data"

func main() {
	planet, result := createWorld("Earth")
	fmt.Printf("Create world: %s\n", result)
	save("Earth", planet.changes)

	savedEarth := load("Earth")
	fmt.Println(savedEarth)
	fmt.Printf("Events stored at %s.\n", eventDataDirectory)
}

func save(planetID string, changes []interface{}) {
	eventStore, err := filestore.FileStore(eventDataDirectory)
	failIf(err)
	for _, event := range changes {
		jsonEvent, err := json.Marshal(event)
		failIf(err)
		eventStore.Save("world", planetID, reflect.TypeOf(event).String(), []byte(jsonEvent))
	}
}

func load(planetID string) *world {
	eventStore, err := filestore.FileStore(eventDataDirectory)
	failIf(err)
	eventHistory, err := eventStore.Get("world", planetID)
	failIf(err)

	var events []interface{}
	for _, item := range eventHistory {
		if item.Type == "main.worldCreated" {
			var event worldCreated
			err := json.Unmarshal(item.Data, &event)
			failIf(err)
			events = append(events, event)
		}
	}

	var planet world
	planet.loadFromHistory(events)
	return &planet
}

func failIf(err error) {
	if err != nil {
		panic(err)
	}
}
