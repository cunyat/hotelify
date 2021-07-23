package get

import (
	"context"
	"errors"

	"github.com/cunyat/hotelify/internal/common/domain/query"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type RoomQuery struct {
	UUID string
}

type RoomQueryResponse struct {
	UUID     string        `json:"uuid"`
	Num      string        `json:"num"`
	Floor    int           `json:"floor"`
	Services []string      `json:"services"`
	Beds     []BedResponse `json:"beds"`
}

type BedResponse struct {
	BedType string `json:"bedType"`
	Count   int    `json:"count"`
}

const RoomQueryType query.Type = "rooms.room.get"

func (q RoomQuery) QueryName() query.Type {
	return RoomQueryType
}

var _ query.Query = (*RoomQuery)(nil)

func RoomQueryHandler(repo room.Repository) query.Handler {
	return func(ctx context.Context, baseQuery query.Query) (query.Response, error) {
		query, ok := baseQuery.(RoomQuery)
		if !ok {
			return RoomQueryResponse{}, errors.New("unknown query")
		}

		room, err := repo.Get(ctx, query.UUID)
		if err != nil {
			return RoomQueryResponse{}, err
		}

		beds := []BedResponse{}

		for _, bed := range room.Beds() {
			bed := BedResponse{
				BedType: bed.String(),
				Count:   bed.Count(),
			}
			beds = append(beds, bed)
		}

		return RoomQueryResponse{
			UUID:     room.UUID(),
			Num:      room.Num(),
			Floor:    room.Floor(),
			Services: room.Services(),
			Beds:     beds,
		}, nil
	}
}
