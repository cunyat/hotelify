package room

import (
	"context"
	"errors"
)

var (
	ErrRoomNotFound = errors.New("room not found")
)

type Repository interface {
	Get(context.Context, string) (Room, error)
	Save(context.Context, Room) error
	List(context.Context) ([]Room, error)
	Update(context.Context, Room) error
	Delete(context.Context, Room) error
}
