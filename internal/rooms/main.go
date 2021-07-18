package main

import (
	"fmt"
	"net/http"

	"github.com/cunyat/hotelify/internal/rooms/adapters/command"
	"github.com/cunyat/hotelify/internal/rooms/adapters/storage"
	"github.com/cunyat/hotelify/internal/rooms/app/create"
	"github.com/cunyat/hotelify/internal/rooms/ports"
	"github.com/go-chi/chi/v5"
)

func main() {
	cbus := command.NewInMemoryCommandBus()
	repo := storage.NewInMemoryRoomRepository()
	createRoom := create.RoomCommandHandler(repo)

	cbus.Register(create.RoomCommand{}.CommandName(), createRoom)

	router := chi.NewRouter()

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	handler := ports.HandlerFromMux(ports.NewHttpServer(&cbus), router)
	rootRouter.Mount("/api", handler)

	fmt.Println("Running")
	http.ListenAndServe(":9050", rootRouter)
}
