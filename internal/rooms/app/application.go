package app

import (
	"github.com/cunyat/hotelify/internal/common/domain/command"
	"github.com/cunyat/hotelify/internal/common/domain/event"
	"github.com/cunyat/hotelify/internal/common/domain/query"
)

type Application struct {
	CommandBus command.Bus
	QueryBus   query.Bus
	EventBus   event.Bus
}
