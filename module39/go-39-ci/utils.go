package main

import (
	"os/user"
)

func getUsername() string {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	return current.Name
}

func sum(a, b int) int {
	return a + b
}
