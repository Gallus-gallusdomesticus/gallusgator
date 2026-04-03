package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command) error {

	if len(cmd.args) < 2 { //handlerFeed need both name and url
		fmt.Println("Command require name and URL.")
		os.Exit(1)
	}

	ctx := context.Background()

	currentuser, err := s.db.GetUser(ctx, s.cfg.CurrentUserName) //get the currently login user data
	if err != nil {
		return err
	}

	feedParam := database.CreateFeedParams{ //make the parameter for the command
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    currentuser.ID,
	}

	feed, err := s.db.CreateFeed(ctx, feedParam)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil

}
