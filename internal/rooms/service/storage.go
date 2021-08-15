package service

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cunyat/hotelify/internal/rooms/adapters/storage"
	"github.com/cunyat/hotelify/internal/rooms/domain/room"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewRoomRepository() room.Repository {
	config, err := loadConfig()
	if err != nil {
		panic(err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/rooms", config.user, config.password, config.host, config.port)
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err == nil {
		panic(err.Error())
	}

	return storage.NewMysqlRoomRepository(pool)
}

func loadConfig() (repoConfig, error) {
	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return repoConfig{}, errors.New("POSTGRES_USER env var not foud")
	}

	passord, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return repoConfig{}, errors.New("POSTGRES_PASSWORD env var not foud")
	}

	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return repoConfig{}, errors.New("POSTGRES_HOST env var not foud")
	}
	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return repoConfig{}, errors.New("POSTGRES_PORT env var not foud")
	}
	db, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return repoConfig{}, errors.New("POSTGRES_DB env var not found")
	}
	return repoConfig{
		user:     user,
		password: passord,
		host:     host,
		port:     port,
		db:       db,
	}, nil
}

type repoConfig struct {
	user     string
	password string
	host     string
	port     string
	db       string
}
