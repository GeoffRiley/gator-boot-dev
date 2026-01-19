package main

import (
	"fmt"

	"github.com/GeoffRiley/gator-boot-dev/internal/database"
)

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
