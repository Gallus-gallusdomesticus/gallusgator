package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error { //login handler function
	if len(cmd.args) == 0 { //expect a single argument
		return fmt.Errorf("Login required a username argument.")

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
		return fmt.Errorf("Register required a username.")

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

func handlerUsers(s *state, cmd command) error { //register users function

	ctx := context.Background() //add context for s.db.GetUsers

	users, err := s.db.GetUsers(ctx) //get users list

	if err != nil {
		return err
	}

	if len(users) == 0 { //check if there is user registered or not
		fmt.Println("No user registered.")
	}

	for _, user := range users {
		if s.cfg.CurrentUserName == user { //if it is the current login user
			fmt.Println("*", user, "(current)")
		} else {
			fmt.Println("*", user) //print the user list
		}

	}

	return nil

}
