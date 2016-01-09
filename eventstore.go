package hist

import (
	"time"
)

// An Eventstore stores and gets events.
type Eventstore interface {

	// Save stores a new event for an instance of an aggregate.
	Save(aggregateType string, aggregateID string, eventType string, eventData []byte) error

	// Get loads all events for an instance of an aggregate.
	Get(aggregateType string, aggregateID string) ([]Event, error)
}

// Event is returned by fileEventstore.Get().
type Event struct {
	// When the event was saved.
	Timestamp time.Time

	// The event type.
	Type string

	// Event data, e.g. serialized JSON
	Data []byte
}
