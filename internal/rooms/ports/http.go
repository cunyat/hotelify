package ports

import (
	"fmt"
	"net/http"

	"github.com/cunyat/hotelify/internal/common/domain"
	"github.com/cunyat/hotelify/internal/rooms/app/create"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type HttpServer struct {
	cBus domain.CommandBus
}

func NewHttpServer(cBus domain.CommandBus) HttpServer {
	return HttpServer{
		cBus: cBus,
	}
}

func (s HttpServer) CreateRoom(w http.ResponseWriter, r *http.Request) {
	postRoom := PostRoom{}

	if err := render.Decode(r, &postRoom); err != nil {
		w.Write([]byte(fmt.Sprintf("{'slug': 'bad-request', 'message': %s}", err.Error())))
		return
	}

	beds := make(map[string]int)
	for _, bed := range *postRoom.Beds {
		bedType := string(*bed.BedType)
		beds[bedType] = *bed.Count
	}

	cmd := create.RoomCommand{
		UUID:     uuid.NewString(),
		Num:      postRoom.Num,
		Floor:    postRoom.Floor,
		Beds:     beds,
		Services: *postRoom.Services,
	}

	if err := s.cBus.Dispatch(r.Context(), cmd); err != nil {
		w.Write([]byte(fmt.Sprintf("{'slug': 'unexpected-error', 'message': %s}", err.Error())))
		return
	}

	w.Header().Set("content-location", "/rooms/"+cmd.UUID)
	w.WriteHeader(http.StatusCreated)
}
