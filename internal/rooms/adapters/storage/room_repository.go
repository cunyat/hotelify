package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/cunyat/hotelify/internal/rooms/domain/room"
)

type InMemoryRoomRepository struct {
	rooms map[string]room.Room
}

var _ room.Repository = (*InMemoryRoomRepository)(nil)

func NewInMemoryRoomRepository() *InMemoryRoomRepository {
	return &InMemoryRoomRepository{
		rooms: make(map[string]room.Room),
	}
}

func (r *InMemoryRoomRepository) Save(ctx context.Context, entity room.Room) error {
	_, ok := r.rooms[entity.UUID()]
	if ok {
		return errors.New("duplicated room uuid")
	}

	r.rooms[entity.UUID()] = entity

	fmt.Printf("New room: %v", entity)
	return nil
}

func (r *InMemoryRoomRepository) Get(ctx context.Context, uuid string) (room.Room, error) {
	entity, ok := r.rooms[uuid]
	if !ok {
		return room.Room{}, room.ErrRoomNotFound
	}

	return entity, nil
}
