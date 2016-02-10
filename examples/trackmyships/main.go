//A more comprehensive domain modeling and eventsourcing example.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/robertreppel/hist"
	"github.com/robertreppel/hist/dynamostore"
	"github.com/robertreppel/hist/examples/trackmyships/ship"
	"github.com/robertreppel/hist/filestore"
)

const useDynamo = true // false = use local file storage.

func main() {
	fmt.Println("\nSHIP LOCATION TRACKER")

	shipName := "Seagull"
	currentLocation := ship.AtSea

	fmt.Printf("\nAdding the '%s' (current location: %s) to the ship's register:\n", shipName, currentLocation)
	result := register(shipName, currentLocation)
	fmt.Printf("Result: %s\n\n", result)

	fmt.Printf("Trying to register the '%s' for the second time:\n", shipName)
	result = register(shipName, currentLocation)
	fmt.Printf("Result: %s\n\n", result)

	fmt.Printf("Recording %s's arrival at the Port of Vancouver:\n", shipName)
	result = recordArrival(shipName, "Vancouver")
	fmt.Printf("Result: %s\n\n", result)

	fmt.Printf("Trying to record %s's arrival in Vancouver for the second time:\n", shipName)
	result = recordArrival(shipName, "Vancouver")
	fmt.Printf("Result: %s\n\n", result)

	fmt.Printf("Departing from Vancouver:\n")
	result = depart(shipName)
	fmt.Printf("Result: %s\n\n", result)

	fmt.Printf("Trying to depart from Vancouver a second time, without first arriving there again:\n")
	result = depart(shipName)
	fmt.Printf("Result: %s\n\n", result)
}

func register(shipName string, currentLocation string) string {
	previousHistory := getPortsOfCallHistory(shipName)

	if shipHas(previousHistory) {
		return "Cannot register ship: It already exists."
	}

	ship, result, err := ship.Register(shipName, ship.AtSea)
	failIf(err)

	updatePortsOfCallHistory(shipName, ship.Changes)
	return result
}

func recordArrival(shipName string, portName string) string {
	ofPastArrivalsAndDepartures := getPortsOfCallHistory(shipName)

	ship, err := ship.NewShipFromHistory(ofPastArrivalsAndDepartures)
	failIf(err)

	result, err := ship.Arrive(portName)
	failIf(err)

	updatePortsOfCallHistory(shipName, ship.Changes)
	return result
}

func depart(shipName string) string {
	ofPastArrivalsAndDepartures := getPortsOfCallHistory(shipName)

	ship, err := ship.NewShipFromHistory(ofPastArrivalsAndDepartures)
	failIf(err)

	result, err := ship.Depart()
	failIf(err)

	updatePortsOfCallHistory(shipName, ship.Changes)
	return result
}

const dataStoreDirectory = "/tmp/hist-example-ship"
const eventstoreTable = "ShipTracker"
const region = "us-west-2"

// If no endpoint is configured, cloud AWS DynamoDB will be used.
// const endpoint = ""

// Local DynamoDB; see https://aws.amazon.com/blogs/aws/dynamodb-local-for-desktop-development/
const endpoint = "http://localhost:8000"

func init() {
	if useDynamo {
		fmt.Println("Using Dynamodb.")
		dynamostore.Create(eventstoreTable, region, endpoint)
	} else {
		fmt.Println("Using Filestore.")
		deleteAllDataFrom(dataStoreDirectory)
	}
}

func getPortsOfCallHistory(shipName string) []interface{} {
	var store hist.Eventstore
	var err error
	if useDynamo {
		store, err = dynamostore.DynamoStore(eventstoreTable, region, endpoint)
		failIf(err)
	} else {
		store, err = filestore.FileStore(dataStoreDirectory)
		failIf(err)
	}

	eventHistory, err := store.Get(shipAggregateType, shipName)
	failIf(err)

	var events []interface{}
	for _, item := range eventHistory {
		if item.Type == "ship.Registered" {
			var event ship.Registered
			err := json.Unmarshal(item.Data, &event)
			failIf(err)
			events = append(events, event)
		}
		if item.Type == "ship.Arrived" {
			var event ship.Arrived
			err := json.Unmarshal(item.Data, &event)
			failIf(err)
			events = append(events, event)
		}
		if item.Type == "ship.Departed" {
			var event ship.Departed
			err := json.Unmarshal(item.Data, &event)
			failIf(err)
			events = append(events, event)
		}
	}
	return events
}

func updatePortsOfCallHistory(shipName string, changes []interface{}) {
	var store hist.Eventstore
	var err error
	if useDynamo {
		store, err = dynamostore.DynamoStore(eventstoreTable, region, endpoint)
		failIf(err)
	} else {
		store, err = filestore.FileStore(dataStoreDirectory)
		failIf(err)
	}
	for _, event := range changes {
		jsonEvent, err := json.Marshal(event)
		failIf(err)
		store.Save("Ship", shipName, reflect.TypeOf(event).String(), []byte(jsonEvent))
	}
}

func shipHas(previousHistory []interface{}) bool {
	return len(previousHistory) > 0
}

func deleteAllDataFrom(dataDirectory string) error {
	err := os.RemoveAll(dataDirectory)
	if err != nil {
		return err
	}
	return nil
}

func failIf(err error) {
	if err != nil {
		panic(err)
	}
}

const shipAggregateType string = "Ship"
