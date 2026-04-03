package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error { //register reset function
	ctx := context.Background() //add context for s.db.ResetUser

	if err := s.db.ResetAll(ctx); err != nil { //resetall function
		fmt.Println("Database reset failed!", err)
		os.Exit(1)
	}

	fmt.Println("Database reset successful.")
	return nil
}
