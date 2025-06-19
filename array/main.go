package main

import (
	"fmt"
)

func main() {
	var users [3]string
	users[0] = "Khadga Bahadur Shrestha"
	users[1] = "Sakar Subedi"
	users[2] = "Himal Ghimire"

	fmt.Println(users)
	fmt.Println(users[0])

	skills := [2]string{
		"flutter", "golang",
	}
	fmt.Println("skills:", skills)

	for i, user := range users {
		fmt.Println(i, user)
	}
}
