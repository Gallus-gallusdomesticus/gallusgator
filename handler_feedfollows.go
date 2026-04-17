package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 { //handlerFeed need url
		return fmt.Errorf("Usage: %s <url>", cmd.name)
	}

	ctx := context.Background()

	feed, err := s.db.GetFeedByURL(ctx, cmd.args[0]) //get the feed by URL
	if err != nil {
		return fmt.Errorf("URL not found: %w", err)
	}

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

	fmt.Printf("Feed Name: %s\n", follow.FeedName) //feed name
	fmt.Printf("User Name: %s\n", follow.UserName) //user that follow

	return nil

}

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	follows, err := s.db.GetFeedFollows(ctx, user.ID) //get the list of follow based on user id
	if err != nil {
		return fmt.Errorf("Current user ID not found: %w", err)
	}

	fmt.Printf("Current user name: %s\n", s.cfg.CurrentUserName)
	fmt.Println("Following:")
	if len(follows) == 0 { //print no feed followed if there is no follow
		fmt.Println("NO FEEDS FOLLOWED!")
	}

	for idx, follow := range follows { //list feeds that is followed
		fmt.Printf("%d. %s\n", idx+1, follow.FeedName)

	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) != 1 { //handlerUnfollow need url
		return fmt.Errorf("Usage: %s <url>", cmd.name)
	}

	ctx := context.Background()

	feed, err := s.db.GetFeedByURL(ctx, cmd.args[0]) //get the feed by URL
	if err != nil {
		return fmt.Errorf("URL not found: %w", err)
	}

	deleteParam := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.db.DeleteFeedFollow(ctx, deleteParam); err != nil {
		return fmt.Errorf("Fail to unfollow: %w", err)
	}

	return nil

}
