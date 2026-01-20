package main

import (
	"context"
	"fmt"
	"github.com/GeoffRiley/gator-boot-dev/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), (user.ID))
	if err != nil {
		return fmt.Errorf("couldn't get feed follows for %s: %w", s.cfg.CurrentUserName, err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No follows found.")
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}

