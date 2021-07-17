package room

import "context"

type Repository interface {
	Get(context.Context, string) (Room, error)
	Save(context.Context, Room) error
	Search(context.Context) []Room
}
