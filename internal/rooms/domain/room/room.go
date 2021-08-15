package room

import (
	"errors"

	uuidgen "github.com/google/uuid"
)

// Room represents a single room in the hotel and it's attributes
type Room struct {
	uuid     string
	num      string
	floor    int
	capacity int
	services []string
}

// NewRoom builds and returns a Room entity
func NewRoom(uuid string, num string, floor int, capacity int, services []string) (Room, error) {
	if uuid == "" {
		return Room{}, errors.New("empty room uuid")
	}

	if capacity <= 0 {
		return Room{}, errors.New("capacity must be a positive integer")
	}

	return Room{
		uuid:     uuid,
		num:      num,
		floor:    floor,
		capacity: capacity,
		services: services,
	}, nil
}

func CreateRoom(uuid string, num string, floor int, capacity int, services []string) (Room, error) {
	room, err := NewRoom(uuid, num, floor, capacity, services)
	if err != nil {
		return Room{}, err
	}

	return room, nil
}

// UUID return room's uuid
func (r Room) UUID() string {
	return r.uuid
}

// Num return room's num
func (r Room) Num() string {
	return r.num
}

// Floor return room's floor
func (r Room) Floor() int {
	return r.floor
}

// Capacity return room's capacity
func (r Room) Capacity() int {
	return r.capacity
}

// Services return room's services
func (r Room) Services() []string {
	return r.services
}

type RoomCreated struct {
	UUID      string
	eventUUID string
}

func NewRoomCreated(uuid string) RoomCreated {
	return RoomCreated{
		UUID:      uuid,
		eventUUID: uuidgen.NewString(),
	}
}

func (rc RoomCreated) EventUUID() string {
	return rc.eventUUID
}

func (rc RoomCreated) AggregateUUID() string {
	return rc.UUID
}

func (rc RoomCreated) EventName() string {
	return "rooms.room_created"
}
