package response

import "github.com/cunyat/hotelify/internal/rooms/domain/room"

type Bed struct {
	BedType string `json:"bedType"`
	Count   int    `json:"count"`
}

type Room struct {
	UUID     string   `json:"uuid"`
	Num      string   `json:"num"`
	Floor    int      `json:"floor"`
	Services []string `json:"services"`
	Beds     []Bed    `json:"beds"`
}

type Rooms struct {
	Rooms []Room `json:"rooms"`
}

func FromDomain(entity room.Room) Room {
	beds := []Bed{}

	for _, bed := range entity.Beds() {
		bed := Bed{
			BedType: bed.String(),
			Count:   bed.Count(),
		}
		beds = append(beds, bed)
	}

	return Room{
		UUID:     entity.UUID(),
		Num:      entity.Num(),
		Floor:    entity.Floor(),
		Services: entity.Services(),
		Beds:     beds,
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
