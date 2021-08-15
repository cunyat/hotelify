package create

import (
	"context"
	"fmt"

	"github.com/cunyat/hotelify/internal/common/domain/command"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type RoomCommand struct {
	UUID     string
	Num      string
	Floor    int
	Capacity int
	Services []string
}

const RoomCommandType command.Type = "rooms.room.create"

func (c RoomCommand) CommandName() command.Type {
	return RoomCommandType
}

var _ command.Command = (*RoomCommand)(nil)

func RoomCommandHandler(repo room.Repository) command.Handler {
	return func(ctx context.Context, baseCmd command.Command) error {
		cmd, ok := baseCmd.(RoomCommand)
		if !ok {
			return fmt.Errorf("unknown command (%s) in create.room.command", baseCmd.CommandName())
		}

		room, err := room.CreateRoom(
			cmd.UUID,
			cmd.Num,
			cmd.Floor,
			cmd.Capacity,
			cmd.Services,
		)
		if err != nil {
			return fmt.Errorf("error creating room: %w", err)
		}

		err = repo.Save(ctx, room)
		if err != nil {
			return fmt.Errorf("error saving room: %w", err)
		}

		return nil
	}
}
