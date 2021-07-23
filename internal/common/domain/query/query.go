package query

import (
	"context"
	"errors"
)

type Type string

type Query interface {
	QueryName() Type
}

type Response interface{}

type Handler func(context.Context, Query) (Response, error)

type Bus interface {
	// Register is the method used to register a new QueryHandler
	Register(Type, Handler)
	// Ask is the method used to ask a new Query and obtain the response
	Ask(context.Context, Query) (Response, error)
}

var (
	ErrQueryNotRegistered = errors.New("query is not registered")
)
