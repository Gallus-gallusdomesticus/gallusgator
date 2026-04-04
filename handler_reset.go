package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error { //register reset function
	ctx := context.Background() //add context for s.db.ResetUser

	if err := s.db.ResetAll(ctx); err != nil { //resetall function
		return fmt.Errorf("Database reset failed! %w", err)

	}

	fmt.Println("Database reset successful.")
	return nil
}
