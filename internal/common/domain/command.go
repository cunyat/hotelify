package domain

import (
	"context"
	"errors"
)

type CommandType string

type Command interface {
	CommandName() CommandType
}

type CommandHandler func(context.Context, Command) error

type CommandBus interface {
	// Register is the method used to register a new command handler.
	Register(CommandType, CommandHandler)
	// Dispatch is the method used to dispatch new commands.
	Dispatch(context.Context, Command) error
}

var ErrCommandNotRegistered = errors.New("command not registered")
