package hist

import (
	"time"
)

// Eventstore stores and gets events.
type Eventstore interface {

	// Save stores a new event for a stream.
	Save(streamID string, eventType string, eventData []byte) error

	// Get gets all events for a stream.
	Get(streamIdID string) ([]Event, error)
}

// Event is returned by fileEventstore.Get().
type Event struct {
	// When the event was saved.
	Timestamp time.Time

	// ID of the stream of data for all events which belong to a transaction.
	StreamID string

	// The event type.
	Type string

	// Event data, e.g. serialized JSON
	Data []byte
}
