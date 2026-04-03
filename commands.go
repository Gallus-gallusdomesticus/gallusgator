package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct { //hold the commands the CLI can handle
	handlers map[string]func(*state, command) error //map of command name (the key is the command name while the value is the handler function)
}

func (c *commands) run(s *state, cmd command) error { //run given command with provided state
	handler, ok := c.handlers[cmd.name]
	if !ok { //check if command exist
		return errors.New("function doesnt exist")
	}

	if err := handler(s, cmd); err != nil { //run the command
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) { //register new command
	c.handlers[name] = f
}
