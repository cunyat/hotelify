package domain

import "context"

type QueryType string

type Query interface {
	QueryName() QueryType
}

type Response interface{}

type QueryHandler func(context.Context, Query) (Response, error)

type QueryBus interface {
	// Register is the method used to register a new QueryHandler
	Register(QueryType, QueryHandler)
	// Ask is the method used to ask a new Query and obtain the response
	Ask(context.Context, Query) (Response, error)
}
