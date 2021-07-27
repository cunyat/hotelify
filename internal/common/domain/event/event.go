package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Type represents a unique identifier for each type of event
type Type string

// Event represents the behaviour of a Domain Event
type Event interface {
	EventUUID() string
	AggregateUUID() string
	OccurredOn() time.Time
	EventType() Type
}

// Listener defines the signature of an event listener
type Listener func(context.Context, Event) error

// Bus defines the behaviour of the Event Bus
type Bus interface {
	// Subscribe register a listener for the given event type
	Subscribe(Type, Listener)

	// Publish sends events to their listeners
	Publish(context.Context, ...Event) error
}

// BaseEvent gives the basic & common features of a domain event
type BaseEvent struct {
	eventUUID     string
	aggregateUUID string
	occurredOn    time.Time
}

// Generates a new BaseEvent for the given aggregate (as uuid)
func NewBaseEvent(aggregateUUID string) BaseEvent {
	return BaseEvent{
		eventUUID:     uuid.NewString(),
		aggregateUUID: aggregateUUID,
		occurredOn:    time.Now(),
	}
}

// EventUUID return the current event identifier
func (b BaseEvent) EventUUID() string {
	return b.eventUUID
}

// OccurredOn represents the moment that the event has been fired
func (b BaseEvent) OccurredOn() time.Time {
	return b.occurredOn
}

// AggregateUUID represents the aggregate that fired a given event
func (b BaseEvent) AggregateUUID() string {
	return b.aggregateUUID
}
