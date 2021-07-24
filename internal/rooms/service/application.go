package service

import (
	"context"

	"github.com/cunyat/hotelify/internal/common/adapters/command"
	"github.com/cunyat/hotelify/internal/common/adapters/query"
	"github.com/cunyat/hotelify/internal/rooms/app"
	"github.com/cunyat/hotelify/internal/rooms/app/create"
	"github.com/cunyat/hotelify/internal/rooms/app/get"
	"github.com/cunyat/hotelify/internal/rooms/app/list"
)

func NewApplication(ctx context.Context) app.Application {
	return newApplication(ctx)
}

func newApplication(ctx context.Context) app.Application {
	cbus := command.NewInMemoryCommandBus()
	qbus := query.NewInMemoryQueryBus()
	repo := NewRoomRepository()

	// Commands
	cbus.Register(create.RoomCommandType, create.RoomCommandHandler(repo))

	// Queries
	qbus.Register(get.RoomQueryType, get.RoomQueryHandler(repo))
	qbus.Register(list.RoomQueryType, list.RoomQueryHandler(repo))

	return app.Application{
		CommandBus: cbus,
		QueryBus:   qbus,
	}
}
