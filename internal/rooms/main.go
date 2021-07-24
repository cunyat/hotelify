package main

import (
	"context"
	"log"
	"os"

	"github.com/cunyat/hotelify/internal/rooms/ports"
	"github.com/cunyat/hotelify/internal/rooms/service"
)

func main() {
	app := service.NewApplication(context.TODO())
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	ctx, srv := ports.NewHttpServer(context.TODO(), port, app)
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
