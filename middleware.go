package main

import (
	"context"
	"fmt"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()
		user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("User not found: %w", err)
		}

		return handler(s, cmd, user)
	}
}
