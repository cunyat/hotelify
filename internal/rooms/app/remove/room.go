package remove

import (
	"context"
	"fmt"

	"github.com/cunyat/hotelify/internal/common/domain/command"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type RoomCommand struct {
	UUID string
}

var RoomCommandType command.Type = "command.room.remove"

func (c RoomCommand) CommandName() command.Type { return RoomCommandType }

var _ command.Command = (*RoomCommand)(nil)

func RoomCommandHandler(repo room.Repository) command.Handler {
	return func(ctx context.Context, bCmd command.Command) error {
		cmd, ok := bCmd.(RoomCommand)
		if !ok {
			return fmt.Errorf("unkown command (%s) in remove.Room", bCmd.CommandName())
		}

		rm, err := repo.Get(ctx, cmd.UUID)
		if err != nil {
			return err
		}

		return repo.Delete(ctx, rm)
	}
}
