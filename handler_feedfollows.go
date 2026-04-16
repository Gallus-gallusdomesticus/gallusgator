package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 { //handlerFeed need url
		return fmt.Errorf("Usage: %s <url>", cmd.name)
	}

	ctx := context.Background()

	feed, err := s.db.GetFeedByURL(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("URL not found: %w", err)
	}

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Current user not found: %w", err)
	}

	followParam := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.db.CreateFeedFollows(ctx, followParam)
	if err != nil {
		return fmt.Errorf("Fail to create feed follows: %w", err)
	}

	fmt.Printf("Feed Name: %s", follow.FeedName)
	fmt.Printf("User Name: %s", follow.UserName)

	return nil

}

func handlerFollowing(s *state, cmd command) error {
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Current user not found: %w", err)
	}

	follows, err := s.db.GetFeedFollows(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("Current user ID not found: %w", err)
	}

	fmt.Printf("Current user name: %s", s.cfg.CurrentUserName)
	fmt.Printf("Following:")
	if len(follows) == 0 {
		fmt.Printf("NO FEEDS FOLLOWED!")
	}

	for idx, follow := range follows {
		fmt.Printf("%d. %s", idx, follow.FeedName)

	}

	return nil
}
