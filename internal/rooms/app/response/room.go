package response

import "github.com/cunyat/hotelify/internal/rooms/domain/room"

type Room struct {
	UUID     string   `json:"uuid"`
	Num      string   `json:"num"`
	Floor    int      `json:"floor"`
	Services []string `json:"services"`
	Capacity int      `json:"capacity"`
}

type Rooms struct {
	Rooms []Room `json:"rooms"`
}

func FromDomain(entity room.Room) Room {
	return Room{
		UUID:     entity.UUID(),
		Num:      entity.Num(),
		Floor:    entity.Floor(),
		Services: entity.Services(),
		Capacity: entity.Capacity(),
	}
}

func RoomsFromDomain(entities []room.Room) Rooms {
	rooms := make([]Room, len(entities))

	for i, entity := range entities {
		rooms[i] = FromDomain(entity)
	}

	return Rooms{
		Rooms: rooms,
	}
}
