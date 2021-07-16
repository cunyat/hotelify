package domain

type DomainEvent interface {
	EventUUID() string
	AggregateUUID() string
	EventName() string
}

type AggregateRoot struct {
	events []DomainEvent
}

func (a *AggregateRoot) Record(events ...DomainEvent) {
	a.events = append(a.events, events...)
}

func (a *AggregateRoot) PullDomainEvents() []DomainEvent {
	events := append([]DomainEvent{}, a.events...)
	a.events = nil
	return events
}

type EventHandler func(DomainEvent)

type EventBus interface {
	Publish([]DomainEvent)
	RegisterHandler(string, EventHandler)
}
