package main

import (
	"fmt"

	"github.com/GeoffRiley/gator-boot-dev/internal/database"
)

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func printCurrentUser(user database.User, s *state) {
	if user.Name == s.cfg.CurrentUserName {
		fmt.Printf(" * %v (current)\n", user.Name)
	} else {
		fmt.Printf(" * %v\n", user.Name)
	}
}
