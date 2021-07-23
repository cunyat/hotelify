package get

import (
	"context"
	"errors"

	"github.com/cunyat/hotelify/internal/common/domain/query"
	"github.com/cunyat/hotelify/internal/rooms/app/response"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type RoomQuery struct {
	UUID string
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
			return response.Room{}, errors.New("unknown query")
		}

		room, err := repo.Get(ctx, query.UUID)
		if err != nil {
			return response.Room{}, err
		}

		return response.FromDomain(room), nil
	}
}
