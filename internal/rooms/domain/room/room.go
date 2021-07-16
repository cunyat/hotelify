package room

import (
	"errors"

	"github.com/cunyat/hotelify/internal/common/domain"
	uuidgen "github.com/google/uuid"
)

// Room represents a single room in the hotel and it's attributes
type Room struct {
	uuid     string
	num      string
	floor    int
	beds     []RoomBed
	services []string
	domain.AggregateRoot
}

// NewRoom builds and returns a Room entity
func NewRoom(uuid string, num string, floor int, beds []RoomBed, services []string) (Room, error) {
	if uuid == "" {
		return Room{}, errors.New("empty room uuid")
	}

	return Room{
		uuid:     uuid,
		num:      num,
		floor:    floor,
		beds:     beds,
		services: services,
	}, nil
}

func CreateRoom(num string, floor int, beds []RoomBed, services []string) (Room, error) {
	uuid := uuidgen.NewString()
	room, err := NewRoom(uuid, num, floor, beds, services)
	if err != nil {
		return Room{}, err
	}

	return room, nil
}

type RoomCreated struct {
	UUID      string
	eventUUID string
}

var roomCreated domain.DomainEvent = RoomCreated{}

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
	capacity := 0

	for _, bed := range r.beds {
		capacity += bed.Capacity() * bed.count
	}

	return capacity
}

// Beds return room's beds
func (r Room) Beds() []RoomBed {
	return r.beds
}

// Services return room's services
func (r Room) Services() []string {
	return r.services
}
