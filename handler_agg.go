package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error { //register agg function

	if len(cmd.args) != 1 { //handlerAgg need time
		return fmt.Errorf("Usage: %s <time_between_reqs>", cmd.name)
	}

	duration, err := time.ParseDuration(cmd.args[0]) //convert the duration string to time.Time
	if err != nil {
		return fmt.Errorf("Time must be a duration string: %w", err)
	}

	log.Printf("Collecting feeds every %s", duration)
	ticker := time.NewTicker(duration) //make a ticker every duration
	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil { //use scrape feed command
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

	postparams := createPostParams(rss.Channel.Item, feed.ID) //make a list of post parameter

	var successpost int                    //make a int that count how many post successfully logged
	for _, postparam := range postparams { //for each post parameter
		_, err := s.db.CreatePost(ctx, postparam) //create a post
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") { //no need to log if the url is duplicate
				continue
			}
			log.Printf("Couldn't create post: %v", err) //will log if there is a non duplicate error midway
			continue
		}

		successpost = successpost + 1 //add to number of logged

	}
	printRSS(rss)
	fmt.Printf("Feed %s collected, %v items found, %v successfully logged\n", feed.Name, len(rss.Channel.Item), successpost)
	return nil

}

func printRSS(RSS *RSSFeed) {
	fmt.Println("+++RSS STRUCT+++++++++++++++++++++++")
	fmt.Println("Title	     :", RSS.Channel.Title)
	fmt.Println("Link	     :", RSS.Channel.Link)
	fmt.Println("Description  :", RSS.Channel.Description)
}

func createPostParams(items []RSSItem, feedID uuid.UUID) []database.CreatePostParams {

	var listPostParam []database.CreatePostParams
	for _, item := range items { //for item in RSSItem

		pubtime, err := timeConvert(item.PubDate) //convert the time
		var publishTime sql.NullTime              //make a sql.NullTime variable
		if err != nil {
			publishTime = sql.NullTime{Valid: false} //if the time is NULL make it valid false
		} else {
			publishTime = sql.NullTime{ //if the time is not NULL
				Time:  pubtime,
				Valid: true,
			}
		}

		description := sql.NullString{ //make a sql.NullString variable
			String: item.Description,       //just put whatever in the description
			Valid:  item.Description != "", //if the description is empty make the valid false (it is NULL)
		}

		PostParam := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishTime,
			FeedID:      feedID,
		}

		listPostParam = append(listPostParam, PostParam) //append the parameter it to the list of parameter

	}
	return listPostParam
}

func timeConvert(s string) (time.Time, error) {
	layouts := []string{ //layout of available time format
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
		time.RFC3339Nano,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"Mon, 2 Jan 2006 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, s) //for each available layout parse it to time.Time
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse published date: %q", s)
}
