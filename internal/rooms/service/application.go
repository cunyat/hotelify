package service

import (
	"context"

	"github.com/cunyat/hotelify/internal/rooms/adapters/command"
	"github.com/cunyat/hotelify/internal/rooms/adapters/storage"
	"github.com/cunyat/hotelify/internal/rooms/app"
	"github.com/cunyat/hotelify/internal/rooms/app/create"
)

func NewApplication(ctx context.Context) app.Application {
	return newApplication(ctx)
}

func newApplication(ctx context.Context) app.Application {
	cbus := command.NewInMemoryCommandBus()
	repo := storage.NewInMemoryRoomRepository()
	createRoom := create.RoomCommandHandler(repo)

	cbus.Register(create.RoomCommand{}.CommandName(), createRoom)

	return app.Application{
		CommandBus: cbus,
	}
}
