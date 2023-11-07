package main

import (
	"time"

	"github.com/Set-Kaung/rssagg/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:user_id`
}

func dbFeedtoFeed(feed database.Feed) *Feed {
	return &Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name:      feed.Name,
		URL:       feed.Url,
		UserID:    feed.UserID,
	}
}

func dbFeedsToFedds(dbFeeds []database.Feed) []*Feed {
	feeds := make([]*Feed, len(dbFeeds), len(dbFeeds))
	for i, f := range dbFeeds {
		feeds[i] = dbFeedtoFeed(f)
	}
	return feeds
}
