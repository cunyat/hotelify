package command

import (
	"context"
	"errors"
)

type Type string

type Command interface {
	CommandName() Type
}

type Handler func(context.Context, Command) error

type Bus interface {
	// Register is the method used to register a new command handler.
	Register(Type, Handler)
	// Dispatch is the method used to dispatch new commands.
	Dispatch(context.Context, Command) error
}

var ErrCommandNotRegistered = errors.New("command not registered")
