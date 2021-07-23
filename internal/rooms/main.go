package main

import (
	"context"
	"log"

	"github.com/cunyat/hotelify/internal/rooms/ports"
	"github.com/cunyat/hotelify/internal/rooms/service"
)

func main() {
	app := service.NewApplication(context.TODO())
	ctx, srv := ports.NewHttpServer(context.TODO(), ":9051", app)
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
