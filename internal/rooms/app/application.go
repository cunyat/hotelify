package app

import "github.com/cunyat/hotelify/internal/common/domain"

type Application struct {
	CommandBus domain.CommandBus
	QueryBus   domain.QueryBus
	EventBus   domain.EventBus
}
