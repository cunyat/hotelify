package create

import (
	"context"
	"errors"
	"fmt"

	"github.com/cunyat/hotelify/internal/common/domain/command"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type RoomCommand struct {
	UUID     string
	Num      string
	Floor    int
	Beds     map[string]int
	Services []string
}

const RoomCommandType command.Type = "rooms.room.create"

func (c RoomCommand) CommandName() command.Type {
	return RoomCommandType
}

var _ command.Command = (*RoomCommand)(nil)

func RoomCommandHandler(repo room.Repository) command.Handler {
	return func(ctx context.Context, cmd command.Command) error {
		createCmd, ok := cmd.(RoomCommand)
		if !ok {
			return errors.New("unknown command")
		}

		room, err := room.CreateRoom(
			createCmd.UUID,
			createCmd.Num,
			createCmd.Floor,
			createCmd.Beds,
			createCmd.Services,
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
