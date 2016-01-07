package main

type world struct {
	name    string
	changes []interface{}
}

func createWorld(name string) (*world, string) {
	if name == "" {
		return nil, "A world must have a name."
	}
	var planet world
	planet.trackChange(worldCreated{Name: name})
	return &planet, "Success"
}

func (planet *world) trackChange(event interface{}) {
	planet.changes = append(planet.changes, event)
	planet.transition(event)
}

func (planet *world) transition(event interface{}) {
	switch e := event.(type) {
	case worldCreated:
		planet.name = e.Name
	}
}

func (planet *world) loadFromHistory(events []interface{}) {
	for _, event := range events {
		planet.transition(event)
	}
}

func (planet *world) String() string {
	return "Hello " + planet.name
}
