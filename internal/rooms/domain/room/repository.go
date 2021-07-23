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
}
