package service

import (
	"context"

	"github.com/cunyat/hotelify/internal/rooms/adapters/command"
	"github.com/cunyat/hotelify/internal/rooms/adapters/query"
	"github.com/cunyat/hotelify/internal/rooms/adapters/storage"
	"github.com/cunyat/hotelify/internal/rooms/app"
	"github.com/cunyat/hotelify/internal/rooms/app/create"
	"github.com/cunyat/hotelify/internal/rooms/app/get"
)

func NewApplication(ctx context.Context) app.Application {
	return newApplication(ctx)
}

func newApplication(ctx context.Context) app.Application {
	cbus := command.NewInMemoryCommandBus()
	qbus := query.NewInMemoryQueryBus()

	repo := storage.NewInMemoryRoomRepository()

	// Commands
	cbus.Register(create.RoomCommandType, create.RoomCommandHandler(repo))

	// Queries
	qbus.Register(get.RoomQueryType, get.RoomQueryHandler(repo))

	return app.Application{
		CommandBus: cbus,
		QueryBus:   qbus,
	}
}
