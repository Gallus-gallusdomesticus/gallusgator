package main

import (
	"context"
	"fmt"
	"os"
)

func handlerAgg(s *state, cmd command) error { //register agg function
	ctx := context.Background()
	feed := "https://www.wagslane.dev/index.xml"

	rss, err := fetchFeed(ctx, feed) //fetchFeed function
	if err != nil {
		fmt.Println("Fetch RSS failed!", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", rss)
	return nil
}
