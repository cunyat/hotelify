package query

import (
	"context"

	"github.com/cunyat/hotelify/internal/common/domain"
)

type InMemoryQueryBus struct {
	handlers map[domain.QueryType]domain.QueryHandler
}

var _ domain.QueryBus = (*InMemoryQueryBus)(nil)

func NewInMemoryQueryBus() *InMemoryQueryBus {
	return &InMemoryQueryBus{
		handlers: make(map[domain.QueryType]domain.QueryHandler),
	}
}

func (b *InMemoryQueryBus) Register(qtype domain.QueryType, handler domain.QueryHandler) {
	b.handlers[qtype] = handler
}

func (b *InMemoryQueryBus) Ask(ctx context.Context, query domain.Query) (domain.Response, error) {
	handler, ok := b.handlers[query.QueryName()]

	if !ok {
		return nil, domain.ErrQueryNotRegistered
	}

	return handler(ctx, query)
}
