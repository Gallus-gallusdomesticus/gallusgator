package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error { //register agg function
	ctx := context.Background()
	feed := "https://www.wagslane.dev/index.xml"

	rss, err := fetchFeed(ctx, feed) //fetchFeed function
	if err != nil {
		return fmt.Errorf("Fetch RSS failed! %w", err)
	}

	fmt.Printf("%+v\n", rss)
	return nil
}
