package handler

import (
	"errors"
	"net/http"

	"github.com/cunyat/hotelify/internal/common/domain/query"
	"github.com/cunyat/hotelify/internal/rooms/app/get"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
	"github.com/gin-gonic/gin"
)

func GetRoomHandler(qbus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuid, ok := ctx.Params.Get("uuid")
		if !ok {
			ctx.JSON(http.StatusBadRequest, "missing uuid")
			return
		}

		query := get.RoomQuery{
			UUID: uuid,
		}

		resp, err := qbus.Ask(ctx, query)
		if err != nil {
			switch true {
			case errors.Is(err, room.ErrRoomNotFound):
				ctx.JSON(http.StatusNotFound, err.Error())
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
