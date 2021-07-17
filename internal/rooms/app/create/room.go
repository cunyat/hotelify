package create

import (
	"context"
	"errors"
	"fmt"

	"github.com/cunyat/hotelify/internal/common/domain"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type RoomCommand struct {
	Num      string
	Floor    int
	Beds     map[string]int
	Services []string
}

const commandType domain.CommandType = "rooms.room.create"

func (c RoomCommand) CommandName() domain.CommandType {
	return commandType
}

var _ domain.Command = (*RoomCommand)(nil)

func RoomCommandHandler(repo room.Repository) domain.CommandHandler {
	return func(ctx context.Context, cmd domain.Command) error {
		createCmd, ok := cmd.(RoomCommand)
		if !ok {
			return errors.New("unknown command")
		}

		var beds = make([]room.RoomBed, len(createCmd.Beds))

		for key, count := range createCmd.Beds {
			bedType, err := room.NewBedTypeFromString(key)
			if err != nil {
				return err
			}

			bed := room.NewRoomBed(bedType, count)
			beds = append(beds, bed)
		}

		room, err := room.CreateRoom(
			createCmd.Num,
			createCmd.Floor,
			beds,
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
