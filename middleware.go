package main

import (
	"context"
	"fmt"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error { //accept handler function with (s,cmd,user) structure and return (s,cmd) function
	return func(s *state, cmd command) error { //return (s ,cmd) function
		ctx := context.Background()
		user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName) //getuser
		if err != nil {
			return fmt.Errorf("User not found: %w", err)
		}

		return handler(s, cmd, user) //return (s, cmd, user) using the user from getuser
	}
}
