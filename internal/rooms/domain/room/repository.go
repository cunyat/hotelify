package room

import "context"

type Repository interface {
	Save(context.Context, Room) error
}
