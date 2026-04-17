package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
)

func handlerAgg(s *state, cmd command) error { //register agg function

	if len(cmd.args) != 1 { //handlerAgg need time
		return fmt.Errorf("Usage: %s <time_between_reqs>", cmd.name)
	}

	duration, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Time must be a duration string: %w", err)
	}

	log.Printf("Collecting feeds every %s", duration)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			return fmt.Errorf("Fail to scrape feed:%w", err)
		}
	}

}

func scrapeFeeds(s *state) error { //register scrapeFeeds function
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx) //get the feed need to be fetched
	if err != nil {
		return fmt.Errorf("Fail to fetch feed: %w", err)
	}

	markfeedParam := database.MarkFeedFetchedParams{ //make the mark parameter
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	}

	if err := s.db.MarkFeedFetched(ctx, markfeedParam); err != nil { //mark the feed that is fetched
		return fmt.Errorf("Fail to mark feed: %w", err)
	}

	rss, err := fetchFeed(ctx, feed.Url) //fetchFeed function
	if err != nil {
		return fmt.Errorf("Fetch RSS failed! %w", err)
	}

	printRSS(rss)

	return nil

}

func printRSS(RSS *RSSFeed) {
	fmt.Println("+++RSS STRUCT+++++++++++++++++++++++")
	fmt.Println("Title	     :", RSS.Channel.Title)
	fmt.Println("Link	     :", RSS.Channel.Link)
	fmt.Println("Description  :", RSS.Channel.Description)
	fmt.Println("Item Title   :")
	if len(RSS.Channel.Item) == 0 {
		fmt.Println("NO ITEM!")
	}
	for _, item := range RSS.Channel.Item {
		fmt.Println("-", item.Title)
	}
}
