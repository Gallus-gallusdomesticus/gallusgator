package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 { //handlerBrowse have optional limit
		return fmt.Errorf("Usage: %s <optional:limit>", cmd.name)
	}

	ctx := context.Background()

	var lim int32
	if len(cmd.args) == 0 {
		lim = 2
	} else {
		n, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("Fail to parse cmd.args:%w", err)
		}
		lim = int32(n)
	}

	getpostsparam := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  lim,
	}

	posts, err := s.db.GetPostsByUser(ctx, getpostsparam)
	if err != nil {
		return fmt.Errorf("Fail to get posts:%w", err)
	}

	for _, post := range posts {
		fmt.Println("-", post.Title)
	}

	return nil
}
