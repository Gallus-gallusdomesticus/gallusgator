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
		return errors.New("Login required a username argument.")

	}

	ctx := context.Background()              //add context for s.db.SetUser
	_, err := s.db.GetUser(ctx, cmd.args[0]) //check if the user exist
	if err != nil {
		fmt.Println(cmd.args[0], "is not exist")
		os.Exit(1)
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil { //use state access to config struct to set username
		return err
	}

	fmt.Println("User has been set.")
	return nil
}

func handlerRegister(s *state, cmd command) error { //register handler function
	if len(cmd.args) == 0 { //expect a single argument
		return errors.New("Register required a username.")

	}

	ctx := context.Background() //add context for s.db.CreateUser

	userParam := database.CreateUserParams{ //user parameter for s.db.CreateUser
		ID:        uuid.New(), //new uuid
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(ctx, userParam) //createuser function to get the user log
	if err != nil {
		fmt.Println("User already exist", err)
		os.Exit(1)
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil { //use state access to config struct to set username
		return err
	}

	fmt.Println("User is successfully created", user) //print the user log
	return nil
}

func handlerReset(s *state, cmd command) error { //register reset function
	ctx := context.Background() //add context for s.db.ResetUser

	if err := s.db.ResetAll(ctx); err != nil {
		fmt.Println("Database reset failed!", err)
		os.Exit(1)
	}

	fmt.Println("Database reset successful.")
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
