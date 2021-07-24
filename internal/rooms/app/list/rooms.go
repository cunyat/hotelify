package list

import (
	"context"
	"errors"
	"fmt"

	"github.com/cunyat/hotelify/internal/common/domain/query"
	"github.com/cunyat/hotelify/internal/rooms/app/response"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type RoomQuery struct {
}

const RoomQueryType query.Type = "rooms.room.list"

func (q RoomQuery) QueryName() query.Type {
	return RoomQueryType
}

var _ query.Query = (*RoomQuery)(nil)

func RoomQueryHandler(repo room.Repository) query.Handler {
	return func(ctx context.Context, baseQuery query.Query) (query.Response, error) {
		_, ok := baseQuery.(RoomQuery)
		if !ok {
			return response.Rooms{}, errors.New("unknown query")
		}

		rooms, err := repo.List(ctx)
		if err != nil {
			return response.Rooms{}, fmt.Errorf("eror listing rooms: %w", err)
		}

		return response.RoomsFromDomain(rooms), nil
	}
}
