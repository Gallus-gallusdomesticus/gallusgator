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

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserFromID(ctx, feed.UserID)
		if err != nil {
			return err
		}

		fmt.Println("==================================")
		fmt.Println("Name	:", user.Name)
		fmt.Println("Feed	:", feed.Name)
		fmt.Println("URL	 :", feed.Url)
	}

	return nil
}
