package ship

import "fmt"

type Ship struct {
	name     string
	location string
	Changes  []interface{}
}

const AtSea string = "At Sea"

func (ship *Ship) transition(event interface{}) {
	switch e := event.(type) {
	case Arrived:
		arrived := Arrived(e)
		ship.location = arrived.Port
	case Registered:
		registered := Registered(e)
		ship.name = registered.ShipID
		ship.location = registered.Location
	case Departed:
		ship.location = AtSea
	}
}

func NewShipFromHistory(events []interface{}) (*Ship, error) {
	ship := &Ship{}
	for _, event := range events {
		ship.transition(event)
	}
	return ship, nil
}

func (ship *Ship) trackChange(event interface{}) {
	ship.transition(event)
	ship.Changes = append(ship.Changes, event)
}

func (ship *Ship) String() string {
	return fmt.Sprintf("Ship: %s\tLocated: %s\n", ship.name, ship.location)
}
