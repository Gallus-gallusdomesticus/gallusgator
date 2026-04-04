package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error { //login handler function
	if len(cmd.args) == 0 { //expect a single argument
		return fmt.Errorf("Usage: %s <name>", cmd.name)

	}

	ctx := context.Background()              //add context for s.db.SetUser
	_, err := s.db.GetUser(ctx, cmd.args[0]) //check if the user exist
	if err != nil {
		return fmt.Errorf("%s is not exist", cmd.args[0])
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil { //use state access to config struct to set username
		return fmt.Errorf("Fail to set user: %w", err)
	}

	fmt.Println("User has been set.")
	return nil
}

func handlerRegister(s *state, cmd command) error { //register handler function
	if len(cmd.args) == 0 { //expect a single argument
		return fmt.Errorf("Usage: %s <username>", cmd.name)

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
		return fmt.Errorf("User already exist: %w", err)
	}

	if err := s.cfg.SetUser(cmd.args[0]); err != nil { //use state access to config struct to set username
		return fmt.Errorf("Fail to set user: %w", err)
	}

	printUser(user)
	fmt.Println("User is successfully created") //print the user log
	return nil
}

func handlerUsers(s *state, cmd command) error { //register users function

	ctx := context.Background() //add context for s.db.GetUsers

	users, err := s.db.GetUsers(ctx) //get users list

	if err != nil {
		return fmt.Errorf("Fail to get users list: %w", err)
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

func printUser(user database.User) {
	fmt.Println("===USER STRUCT=======================")
	fmt.Println("ID			:", user.ID)
	fmt.Println("Name		  :", user.Name)
	fmt.Println("Created at	:", user.CreatedAt)
	fmt.Println("Updated at    :", user.UpdatedAt)
}
