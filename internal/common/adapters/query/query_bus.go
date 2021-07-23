package query

import (
	"context"

	"github.com/cunyat/hotelify/internal/common/domain/query"
)

type InMemoryQueryBus struct {
	handlers map[query.Type]query.Handler
}

var _ query.Bus = (*InMemoryQueryBus)(nil)

func NewInMemoryQueryBus() *InMemoryQueryBus {
	return &InMemoryQueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

func (b *InMemoryQueryBus) Register(qtype query.Type, handler query.Handler) {
	b.handlers[qtype] = handler
}

func (b *InMemoryQueryBus) Ask(ctx context.Context, q query.Query) (query.Response, error) {
	handler, ok := b.handlers[q.QueryName()]

	if !ok {
		return nil, query.ErrQueryNotRegistered
	}

	return handler(ctx, q)
}
