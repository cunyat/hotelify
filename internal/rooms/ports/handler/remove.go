package handler

import (
	"errors"
	"net/http"

	"github.com/cunyat/hotelify/internal/common/domain/command"
	"github.com/cunyat/hotelify/internal/rooms/app/remove"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
	"github.com/gin-gonic/gin"
)

func RemoveRoomHandler(cbus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid, ok := ctx.Params.Get("uuid")
		if !ok {
			ctx.JSON(http.StatusBadRequest, "missing uuid")
			return
		}

		cmd := remove.RoomCommand{UUID: uuid}
		err := cbus.Dispatch(ctx, cmd)
		if err != nil {
			switch true {
			case errors.Is(err, room.ErrRoomNotFound):
				ctx.JSON(http.StatusNotFound, err.Error())
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}

		ctx.Status(http.StatusOK)
	}
}
