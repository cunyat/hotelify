package create

import (
	"context"
	"errors"
	"fmt"

	"github.com/cunyat/hotelify/internal/common/domain"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type RoomCommand struct {
	UUID     string
	Num      string
	Floor    int
	Beds     map[string]int
	Services []string
}

const RoomCommandType domain.CommandType = "rooms.room.create"

func (c RoomCommand) CommandName() domain.CommandType {
	return RoomCommandType
}

var _ domain.Command = (*RoomCommand)(nil)

func RoomCommandHandler(repo room.Repository) domain.CommandHandler {
	return func(ctx context.Context, cmd domain.Command) error {
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
