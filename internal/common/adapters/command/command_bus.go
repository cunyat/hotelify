package command

import (
	"context"

	"github.com/cunyat/hotelify/internal/common/domain/command"
)

type InMemoryCommandBus struct {
	handlers map[command.Type]command.Handler
}

var _ command.Bus = (*InMemoryCommandBus)(nil)

func NewInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (b *InMemoryCommandBus) Register(ctype command.Type, handler command.Handler) {
	b.handlers[ctype] = handler
}

func (b *InMemoryCommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handlers[cmd.CommandName()]

	if !ok {
		return command.ErrCommandNotRegistered
	}

	return handler(ctx, cmd)
}
