package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	var char *RSSFeed
	if err := xml.Unmarshal(body, &char); err != nil {
		return nil, err
	}

	char.Channel.Title = html.UnescapeString(char.Channel.Title)
	char.Channel.Description = html.UnescapeString(char.Channel.Description)

	for idx, _ := range char.Channel.Item {
		char.Channel.Item[idx].Title = html.UnescapeString(char.Channel.Item[idx].Title)
		char.Channel.Item[idx].Description = html.UnescapeString(char.Channel.Item[idx].Description)
	}

	return char, nil

}
