package ports

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cunyat/hotelify/internal/rooms/app"
	"github.com/cunyat/hotelify/internal/rooms/ports/handler"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	app app.Application
}

func NewHttpServer(ctx context.Context, httpAddr string, app app.Application) (context.Context, HttpServer) {
	srv := HttpServer{
		httpAddr: httpAddr,
		engine:   gin.New(),

		shutdownTimeout: 15 * time.Second,

		app: app,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *HttpServer) registerRoutes() {
	s.engine.Use(gin.Recovery(), gin.Logger())

	s.engine.POST("/rooms", handler.CreateRoomHandler(s.app.CommandBus))
}

func (s *HttpServer) Run(ctx context.Context) error {
	log.Println("Server running on ", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
