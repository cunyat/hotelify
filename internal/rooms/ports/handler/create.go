package handler

import (
	"net/http"

	"github.com/cunyat/hotelify/internal/common/domain/command"
	"github.com/cunyat/hotelify/internal/rooms/app/create"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createRoomRequest struct {
	Num      string   `json:"num"`
	Floor    int      `json:"floor"`
	Services []string `json:"services"`
	Beds     []struct {
		BedType string `json:"bedType"`
		Count   int    `json:"count"`
	} `json:"beds"`
}

func CreateRoomHandler(cbus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRoomRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		beds := make(map[string]int)
		for _, bed := range req.Beds {
			bedType := string(bed.BedType)
			beds[bedType] = bed.Count
		}

		cmd := create.RoomCommand{
			UUID:     uuid.NewString(),
			Num:      req.Num,
			Floor:    req.Floor,
			Services: req.Services,
			Beds:     beds,
		}

		err := cbus.Dispatch(ctx, cmd)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
