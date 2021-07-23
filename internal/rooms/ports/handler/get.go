package handler

import (
	"net/http"

	"github.com/cunyat/hotelify/internal/common/domain"
	"github.com/cunyat/hotelify/internal/rooms/app/get"
	"github.com/gin-gonic/gin"
)

func GetRoomHandler(qbus domain.QueryBus) gin.HandlerFunc {
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
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
