package command

import (
	"context"

	"github.com/cunyat/hotelify/internal/common/domain"
)

type InMemoryCommandBus struct {
	handlers map[domain.CommandType]domain.CommandHandler
}

var _ domain.CommandBus = (*InMemoryCommandBus)(nil)

func NewInMemoryCommandBus() InMemoryCommandBus {
	return InMemoryCommandBus{
		handlers: make(map[domain.CommandType]domain.CommandHandler),
	}
}

func (b *InMemoryCommandBus) Register(ctype domain.CommandType, handler domain.CommandHandler) {
	b.handlers[ctype] = handler
}

func (b *InMemoryCommandBus) Dispatch(ctx context.Context, cmd domain.Command) error {
	handler, ok := b.handlers[cmd.CommandName()]

	if !ok {
		return domain.ErrCommandNotRegistered
	}

	return handler(ctx, cmd)
}
