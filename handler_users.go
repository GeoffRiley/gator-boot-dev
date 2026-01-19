package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get list of users: <%s>%w", cmd.Name, err)
	}

	for _, user := range users {
		printCurrentUser(user, s)
	}

	return nil
}
