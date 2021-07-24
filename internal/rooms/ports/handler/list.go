package handler

import (
	"net/http"

	"github.com/cunyat/hotelify/internal/common/domain/query"
	"github.com/cunyat/hotelify/internal/rooms/app/list"
	"github.com/gin-gonic/gin"
)

func ListRoomsHandler(qbus query.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		qry := list.RoomQuery{}
		resp, err := qbus.Ask(ctx, qry)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
