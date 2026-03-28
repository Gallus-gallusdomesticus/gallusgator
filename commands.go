package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
	"github.com/google/uuid"
)

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error { //login handler function
	if len(cmd.args) == 0 { //expect a single argument
		return errors.New("login required a username argument")

	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil { //use state access to config struct to set username
		return err
	}

	fmt.Printf("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 { //expect a single argument
		return errors.New("register required a username")

	}

	ctx := context.Background()

	userParam := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(ctx, userParam)
	if err != nil {
		fmt.Println("User already exist", err)
		os.Exit(1)
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil { //use state access to config struct to set username
		return err
	}

	fmt.Println("User is successfully created", user)
	return nil
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
