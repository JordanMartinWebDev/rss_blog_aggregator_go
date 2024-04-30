package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jordanmatinwebdev/rss_blog_agregator_go/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
}

type Feed_Follow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		ApiKey:    user.ApiKey,
	}
}

func databaseFeedToFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		UserID:    feed.UserID,
		Name:      feed.Name,
		Url:       feed.Url,
	}
}

func databaseFeedsToFeeds(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))
	for i, feed := range feeds {
		result[i] = databaseFeedToFeed(feed)
	}
	return result
}

func databaseFeedFollowToFeedFollow(feed_follow database.FeedFollow) Feed_Follow {
	return Feed_Follow{
		ID:        feed_follow.FeedID,
		CreatedAt: feed_follow.CreatedAt,
		UpdatedAt: feed_follow.UpdatedAt,
		UserID:    feed_follow.UserID,
		FeedID:    feed_follow.FeedID,
	}
}
