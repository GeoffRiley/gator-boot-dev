package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/GeoffRiley/gator-boot-dev/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <URL>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't get feed for %s from db: %w", url, err)
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("unable to create feed follow: %w", err)
	}

	fmt.Println("New follow created:")
	printFeedFollow(feed_follow.UserName, feed_follow.FeedName)

	return nil
}

