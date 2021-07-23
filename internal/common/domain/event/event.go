package event

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Type string

type Event interface {
	UUID() string
	AggregateUUID() string
	OccurredOn() time.Time
	EventType() Type
}

type EventListener func(context.Context, Event) error

type Bus interface {
	Subscribe(Type, EventListener)
	Publish(context.Context, Event) error
}

type BaseEvent struct {
	eventUUID     string
	aggregateUUID string
	occurredOn    time.Time
}

func NewBaseEvent(aggregateUUID string) BaseEvent {
	return BaseEvent{
		eventUUID:     uuid.NewString(),
		aggregateUUID: aggregateUUID,
		occurredOn:    time.Now(),
	}
}

func (b BaseEvent) UUID() string {
	return b.eventUUID
}

func (b BaseEvent) OccurredOn() time.Time {
	return b.occurredOn
}

func (b BaseEvent) AggregateUUID() string {
	return b.aggregateUUID
}
