package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
	"github.com/google/uuid"
)

func handlerFeed(s *state, cmd command, user database.User) error {

	if len(cmd.args) < 2 { //handlerFeed need both name and url
		return fmt.Errorf("Usage: %s <name> <url>", cmd.name)
	}

	ctx := context.Background()

	feedParam := database.CreateFeedParams{ //make the parameter for the command
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, feedParam)
	if err != nil {
		return fmt.Errorf("Fail to create feed: %w", err)
	}

	printFeed(feed, user)
	fmt.Println("Feed successfully added.")

	followParam := database.CreateFeedFollowsParams{ //create parameter for feed follow
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.db.CreateFeedFollows(ctx, followParam) //use feed follow command
	if err != nil {
		return fmt.Errorf("Fail to create feed follows: %w", err)
	}

	fmt.Printf("%s successfully added to follow\n", follow.FeedName)

	return nil

}

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("Failed to get feeds: %w", err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserFromID(ctx, feed.UserID)
		if err != nil {
			return fmt.Errorf("Failed to get user: %w", err)
		}

		printFeed(feed, user)
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Println("+++FEED STRUCT+++++++++++++++++++++++")
	fmt.Println("ID			:", feed.ID)
	fmt.Println("Feed		  :", feed.Name)
	fmt.Println("URL		   :", feed.Url)
	fmt.Println("Created at	:", feed.CreatedAt)
	fmt.Println("Updated at    :", feed.UpdatedAt)
	fmt.Println("User		  :", user.Name)
}
